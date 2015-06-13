package engine

import "net/http"

type (

	writerInterface interface {

		http.ResponseWriter
		Done() bool
		Size() int
		Code() int
	}

	writer struct {

		http.ResponseWriter
		size int
		code int
	}
)

func createWriter(w http.ResponseWriter) writerInterface {
	return &writer{w, 0, 0}
}

func (w *writer) Done() bool {
	return w.size != -1
}

func (w *writer) Code() int {
	return w.code
}

func (w *writer) Size() int {
	return w.size
}

func (w *writer) Write(d []byte) (int, error) {

	w.ResponseWriter.WriteHeader(w.code)
	size, err := w.ResponseWriter.Write(d)
	w.size += size

	return size, err
}

func (w *writer) WriteHeader(s int) {
	w.code = s
}
