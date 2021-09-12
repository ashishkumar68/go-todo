package response

import "net/http"

func SendBadRequestContentError(w http.ResponseWriter, message []byte) error {
	w.WriteHeader(400)
	_, err := w.Write(message)
	if err != nil {
		return err
	}

	return nil
}

func SendInternalServerError(w http.ResponseWriter) error {
	w.WriteHeader(500)
	_, err := w.Write([]byte("Internal server error occurred."))
	if err != nil {
		return err
	}

	return nil
}
