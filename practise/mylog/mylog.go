package mylog

import (
	"log"
	"os"
	"fmt"
	"log/syslog"
)

func Sample_log() {
	// 初始化日志记录器
	info_logger := log.New(os.Stdout, "INFO: ", log.LstdFlags | log.Lshortfile)
	debug_logger := log.New(os.Stdout, "DEBUG: ", log.LstdFlags | log.Lshortfile)
	// 使用日志记录器
	info_logger.Output(1, "this is a info message")
	debug_logger.Output(1, "this is a debug message")
	// 隐藏debug日志
	// debug_logger.SetOutput()
	// debug_logger.Output(1, "this is a debug message after set os.DevNull")
}

func Sample_syslog() {
	info_kernel_logger, err := syslog.Dial("", "", syslog.LOG_INFO | syslog.LOG_KERN, "kernel_info")
	if err != nil {
		log.Fatal(err)
	}
	info_deamon_logger, err := syslog.Dial("", "", syslog.LOG_INFO | syslog.LOG_DAEMON, "daemon_info")
	if err != nil {
		log.Fatal(err)
	}
	warn_kernel_logger, err := syslog.Dial("", "", syslog.LOG_WARNING | syslog.LOG_KERN, "kernel_warning")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(info_kernel_logger, "a log message from info_kernel_logger")
	fmt.Fprintf(info_deamon_logger, "a log message from info_deamon_logger")
	// in mac, only warning severity can be record in system.log 
	fmt.Fprintf(warn_kernel_logger, "a log message from warn_kernel_logger")
}