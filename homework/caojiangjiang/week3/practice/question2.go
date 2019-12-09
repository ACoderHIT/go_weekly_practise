package main

/**
 * 方法功能：平滑重启
 */
import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var (
	server   *http.Server
	listener net.Listener
	// 在 go 标准库中提供了一个包：flag，方便进行命令行解析
	// flag.Xxx()，其中 Xxx 可以是 Int、String，Bool 等；返回一个相应类型的指针
	graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
)

func sleep(w http.ResponseWriter, r *http.Request) {
	// 请输入(x)s，比如：10s，不要忘记加s
	duration, err := time.ParseDuration(r.FormValue("duration"))
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	time.Sleep(duration)
	w.Write([]byte("Hello World"))
}


//	父进程监听重启信号
//	在收到重启信号后，父进程调用 fork ，同时传递 socket 描述符给子进程
//	子进程接收并监听父进程传递的 socket 描述符
//	在子进程启动成功之后，父进程停止接收新连接，同时等待旧连接处理完成（或超时）
//	父进程退出，热重启完成

// 执行步骤：
//     1. 启动这个程序（你要是不知道咋启动，劝你还是别看这代码了，看了你也看不懂，懂了你也不会用）
//     2. 浏览器运行localhost:5007/sleep
//     3. 打开iterm 输入 kill -USR2 进程号（用ps -ef|grep question就可以搜到了）
//     4. 观察输出
func main() {
	// 解析命令行参数到定义的flag
	flag.Parse()

	http.HandleFunc("/sleep", sleep)
	server = &http.Server{Addr: ":5007"}

	var err error
	if *graceful {
		log.Print("main: Listening to existing file descriptor 3.")
		// cmd.ExtraFiles: If non-nil, entry i becomes file descriptor 3+i.
		// when we put socket FD at the first entry, it will always be 3(0+3)
		// 子进程的 0, 1, 2 是预留给标准输入、标准输出、错误输出，故传递的 socket 描述符应放在子进程的 3
		f := os.NewFile(3, "")
		listener, err = net.FileListener(f)
	} else {
		log.Print("main: Listening on a new file descriptor.")
		listener, err = net.Listen("tcp", server.Addr)
	}

	if err != nil {
		log.Fatalf("listener error: %v", err)
	}

	go func() {
		// server.Shutdown() stops Serve() immediately, thus server.Serve() should not be in main goroutine
		err = server.Serve(listener)
		log.Printf("server.Serve err: %v\n", err)
	}()
	signalHandler()
	log.Printf("signal end")
}

func reload() error {
	tl, ok := listener.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}
	// 设置传递给子进程的参数（包含 socket 描述符）
	args := []string{"-graceful=true"}
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// put socket FD at the first entry
	cmd.ExtraFiles = []*os.File{f}
	log.Println(cmd)
	return cmd.Start()
}

func signalHandler() {
	ch := make(chan os.Signal, 1)
	// 注册一个os.Signal以监听
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)
	for {
		sig := <-ch
		log.Printf("signal: %v", sig)

		// timeout context for shutdown
		ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			// stop
			log.Printf("stop")
			signal.Stop(ch)
			server.Shutdown(ctx)
			log.Printf("graceful shutdown")
			return
		case syscall.SIGUSR2:
			// reload
			log.Printf("reload")
			err := reload()
			if err != nil {
				log.Fatalf("graceful restart error: %v", err)
			}
			server.Shutdown(ctx)
			log.Printf("graceful reload")
			return
		}
	}
}