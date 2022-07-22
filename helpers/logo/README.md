# Wrapper aroung logurus which allows you to set entry without need to maintain the global variable

## Before

```golang
package logA
var LogrusEntry *logrus.Entry
func init() {
	LogrusEntry = logrus.New().WithFields(logrus.Fields{"app","test"})
}
```

```golang
package main
import xyz.com/xyz/logA
func main() {
	logA.LogrusEntry.Info("Hello world")
}
```

## After

```golang
package logA
import github.com/MyriadFlow/cosmos-wallet/helpers/logo
func init() {
	logrusEntry = logrus.New().WithFields(logrus.Fields{"app","test"})
    logo.SetInstance(logrusEntry)
}
```

```golang
package main
import github.com/MyriadFlow/cosmos-wallet/helpers/logo
func main() {
	logo.Info("Hello world")
}
```
