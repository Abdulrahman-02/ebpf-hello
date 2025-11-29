# eBPF Hello World: Syscall Tracer

A simple eBPF program that traces `execve` syscalls (command executions) on Linux. This project serves as an introduction to eBPF concepts using Go.

## 1. What is eBPF? (The Simple Version)

Imagine the Linux Kernel (the core of the OS) is a highly secure building. Normally, changing how it works is difficult and risky.

**eBPF is like a visitor pass.** It allows running small, safe programs *inside* the kernel to observe system behavior or modify network traffic, without needing to rebuild the kernel or load dangerous modules.

## 2. Project Structure

### 🕵️ `main.bpf.c` (The Spy)
The **C code** that runs inside the kernel.
-   **Role**: It waits for the `sys_enter_execve` event (triggered when a program starts) and logs a message.
-   **Analogy**: A security camera installed in the kernel's hallway.

### 👩‍✈️ `main.go` (The Manager)
The **Go code** that runs in userspace (normal terminal).
-   **Role**: It compiles the C code, loads it into the kernel, and manages its lifecycle.
-   **Analogy**: The security guard monitoring the camera feed.

### 📖 `vmlinux.h` (The Dictionary)
A generated file containing kernel type definitions, allowing the C code to understand kernel structures.

### 🌉 `bpf_x86_bpfel.go` (The Bridge)
Auto-generated Go code that wraps the C program, enabling communication between Go and the kernel.

## 3. Prerequisites

### Install Tools (Ubuntu/Debian)
Required compilers and kernel headers:
```bash
sudo apt update
sudo apt install -y clang llvm libbpf-dev linux-tools-$(uname -r) linux-tools-common linux-headers-$(uname -r)
```

### Generate `vmlinux.h`
Required for kernel types:
```bash
bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
```

## 4. Usage

This project uses [Task](https://taskfile.dev/) for automation.

### Build
Compile the eBPF C code and the Go binary:
```bash
task
```

### Run
Run the tracer (requires sudo):
```bash
task run
```

### View Output
In a separate terminal, read the kernel trace pipe:
```bash
sudo cat /sys/kernel/debug/tracing/trace_pipe
```

### Development
Format code and run linters:
```bash
task fmt
task lint
```

## 5. CI/CD

The project includes a GitHub Actions workflow (`.github/workflows/ci.yml`) that automatically:
-   Installs dependencies.
-   Compiles the code.
-   Runs `golangci-lint` to ensure code quality.
