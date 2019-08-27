# myerrors

Package which helps me to deal with errors:
- Wraps **github.com/pkg/errors** package to support new interfaces in go 1.13 **errors**
- Adds **func StackTrace(error)**
- Adds **Handler** interface which is used in my packages
- Adds **func Merge(error, error)**
