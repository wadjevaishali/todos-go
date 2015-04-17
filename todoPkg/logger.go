package todoPkg

import (
	"log"
	"net/http"
	"time"
)

// var (
// 	Trace   *log.Logger
// 	Info    *log.Logger
// 	Warning *log.Logger
// 	Error   *log.Logger
// )
//
// func Init(
// 	traceHandle io.Writer,
// 	infoHandle io.Writer,
// 	warningHandle io.Writer,
// 	errorHandle io.Writer) {
//
// 	Trace = log.New(traceHandle,
// 		"TRACE: ",
// 		log.Ldate|log.Ltime|log.Lshortfile)
//
// 	Info = log.New(infoHandle,
// 		"INFO: ",
// 		log.Ldate|log.Ltime|log.Lshortfile)
//
// 	Warning = log.New(warningHandle,
// 		"WARNING: ",
// 		log.Ldate|log.Ltime|log.Lshortfile)
//
// 	Error = log.New(errorHandle,
// 		"ERROR: ",
// 		log.Ldate|log.Ltime|log.Lshortfile)
// }

func Logger(inner http.Handler, name string) http.Handler {
	//Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	// Info.Printf("%s\t%s\t%s\t%s",
	//   r.Method,
	//   r.RequestURI,
	//   name,
	//   time.Since(start),
	// )
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf("%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
