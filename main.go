package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	// Remove resource limits for kernels < 5.11
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}

	// Load pre-compiled programs into the kernel.
	objs := bpfObjects{}
	if err := loadBpfObjects(&objs, nil); err != nil {
		log.Fatalf("loading objects: %v", err)
	}
	defer func() {
		if err := objs.Close(); err != nil {
			log.Printf("Error closing objects: %v", err)
		}
	}()

	// Attach to the tracepoint.
	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.TraceExecve, nil)
	if err != nil {
		log.Fatalf("opening tracepoint: %v", err)
	}
	defer func() {
		if err := tp.Close(); err != nil {
			log.Printf("Error closing tracepoint: %v", err)
		}
	}()

	log.Println("eBPF program loaded and attached! Check /sys/kernel/debug/tracing/trace_pipe to see output.")
	log.Println("Press Ctrl+C to exit.")

	// Wait for a signal to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
