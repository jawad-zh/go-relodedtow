package goreloaded


func Capitalize(s string) string {
    str := []rune(s)
    for i := 0; i < len(s); i++ {
        if 'A' <= str[i] && str[i] <= 'Z' {
            str[i] = str[i] + 32
        }
    }
    for i := 0; i < len(str); i++ {
        if (i == 0) && ('a' <= str[i]) && (str[i] <= 'z') {
            str[i] = str[i] - 32
        } else if (str[i] >= 'a' && str[i] <= 'z') && !((str[i-1] >= 'a' && str[i-1] <= 'z') || (str[i-1] >= 'A' && str[i-1] <= 'Z') || (str[i-1] >= '0' && str[i-1] <= '9')) {
            str[i] = str[i] - 32
        } else if ('a' <= str[i]) && (str[i] <= 'z') && (str[i-1] == ' ') {
            str[i] = str[i] - 32
        }
    }
    return string(str)
}