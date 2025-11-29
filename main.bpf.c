//go:build ignore

#include "vmlinux.h"
#include <bpf/bpf_helpers.h>

char __license[] SEC("license") = "Dual MIT/GPL";

SEC("tracepoint/syscalls/sys_enter_execve")
int trace_execve(void *ctx) {
    bpf_printk("Hello from eBPF! Executing a command.\n");
    return 0;
}
