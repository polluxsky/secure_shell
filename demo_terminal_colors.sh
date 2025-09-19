#!/bin/bash

# 定义ANSI颜色代码
RESET="\033[0m"
BOLD="\033[1m"
DIM="\033[2m"
ITALIC="\033[3m"
UNDERLINE="\033[4m"

# 前景色
FG_BLACK="\033[30m"
FG_RED="\033[31m"
FG_GREEN="\033[32m"
FG_YELLOW="\033[33m"
FG_BLUE="\033[34m"
FG_MAGENTA="\033[35m"
FG_CYAN="\033[36m"
FG_WHITE="\033[37m"
FG_BRIGHT_BLACK="\033[90m"
FG_BRIGHT_RED="\033[91m"
FG_BRIGHT_GREEN="\033[92m"
FG_BRIGHT_YELLOW="\033[93m"
FG_BRIGHT_BLUE="\033[94m"
FG_BRIGHT_MAGENTA="\033[95m"
FG_BRIGHT_CYAN="\033[96m"
FG_BRIGHT_WHITE="\033[97m"

# 背景色
BG_BLACK="\033[40m"
BG_RED="\033[41m"
BG_GREEN="\033[42m"
BG_YELLOW="\033[43m"
BG_BLUE="\033[44m"
BG_MAGENTA="\033[45m"
BG_CYAN="\033[46m"
BG_WHITE="\033[47m"

# 打印带颜色的文本的函数
print_color() {
    local text="$1"
    local style="$2"
    local fg="$3"
    local bg="$4"
    echo -ne "${style}${fg}${bg}${text}${RESET}"
}

println_color() {
    print_color "$1" "$2" "$3" "$4"
    echo
}

# 清屏
clear

# 欢迎信息
println_color "=== 终端颜色和命令行风格演示 ===" "${BOLD}" "${FG_CYAN}" "${BG_BLACK}"
echo

# 模拟不同的命令行提示符样式
echo "命令行提示符样式模拟："
echo

# 模拟bash/zsh风格的提示符
print_color "user@machine" "${BOLD}" "${FG_GREEN}" "${BG_BLACK}"
print_color ":" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"
print_color "~/project" "${BOLD}" "${FG_BLUE}" "${BG_BLACK}"
print_color "$ " "${BOLD}" "${FG_WHITE}" "${BG_BLACK}"
println_color "ls -la" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"

# 模拟命令输出
sleep 0.5
println_color "total 32" "${RESET}" "${FG_BRIGHT_BLUE}" "${BG_BLACK}"
println_color "drwxr-xr-x  5 user  staff  160 Sep 19 18:00 ." "${RESET}" "${FG_BRIGHT_BLUE}" "${BG_BLACK}"
println_color "drwxr-xr-x  3 user  staff   96 Sep 19 17:00 .." "${RESET}" "${FG_BRIGHT_BLUE}" "${BG_BLACK}"
print_color "-rw-r--r--  1 user  staff  512 Sep 19 17:50 " "${RESET}" "${FG_BRIGHT_BLUE}" "${BG_BLACK}"
println_color "main.go" "${RESET}" "${FG_GREEN}" "${BG_BLACK}"
print_color "-rwxr-xr-x  1 user  staff 2048 Sep 19 18:00 " "${RESET}" "${FG_BRIGHT_BLUE}" "${BG_BLACK}"
println_color "color-demo" "${RESET}" "${FG_RED}" "${BG_BLACK}"
echo

# 模拟root提示符
print_color "root@machine" "${BOLD}" "${FG_RED}" "${BG_BLACK}"
print_color ":" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"
print_color "/root" "${BOLD}" "${FG_BLUE}" "${BG_BLACK}"
print_color "# " "${BOLD}" "${FG_WHITE}" "${BG_BLACK}"
println_color "sudo systemctl status nginx" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"
echo

# 模拟Kubernetes命令行
print_color "kubectl get pods" "${BOLD}" "${FG_YELLOW}" "${BG_BLACK}"
echo
sleep 0.5
print_color "NAME" "${BOLD}" "${FG_WHITE}" "${BG_BLUE}"
print_color "         READY   STATUS    RESTARTS   AGE" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"
echo
print_color "web-1" "${RESET}" "${FG_GREEN}" "${BG_BLACK}"
print_color "       1/1     Running   0          3d" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"
echo
print_color "web-2" "${RESET}" "${FG_GREEN}" "${BG_BLACK}"
print_color "       1/1     Running   0          3d" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"
echo
print_color "db-1" "${RESET}" "${FG_GREEN}" "${BG_BLACK}"
print_color "        1/1     Running   0          5d" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"
echo

# 模拟Python解释器提示符
print_color "Python 3.9.6 (default, Jun 29 2021, 00:00:00)" "${RESET}" "${FG_YELLOW}" "${BG_BLACK}"
echo
print_color "[GCC 11.1.0] on linux" "${RESET}" "${FG_YELLOW}" "${BG_BLACK}"
echo
print_color "Type 'help', 'copyright', 'credits' or 'license' for more information." "${RESET}" "${FG_YELLOW}" "${BG_BLACK}"
echo
print_color ">>> " "${RESET}" "${FG_GREEN}" "${BG_BLACK}"
println_color "print('Hello, World!')" "${RESET}" "${FG_WHITE}" "${BG_BLACK}"
sleep 0.5
println_color "Hello, World!" "${RESET}" "${FG_CYAN}" "${BG_BLACK}"
print_color ">>> " "${RESET}" "${FG_GREEN}" "${BG_BLACK}"
echo

echo
println_color "=== 演示结束 ===" "${BOLD}" "${FG_CYAN}" "${BG_BLACK}"