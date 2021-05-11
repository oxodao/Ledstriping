package routes

import (
	"encoding/json"
	"fmt"
	"github.com/oxodao/ledstrip/services"
	"net/http"
)

func state(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state, err := prv.ExecuteCommand("state")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write([]byte(state))
	}
}

func debug(prv *services.Provider) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s1, err := prv.ExecuteCommand("dbg")
		fmt.Println("s1", s1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		s2 := prv.ReadLine()
		fmt.Println("s2", s2)

		var obj map[string]json.RawMessage
		err = json.Unmarshal([]byte(s1), &obj)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		var obj2 map[string]json.RawMessage
		err = json.Unmarshal([]byte(s2), &obj2)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		str, _ := json.Marshal(struct {
			PreviousState map[string]json.RawMessage
			CurrentState map[string]json.RawMessage
		}{
			PreviousState: obj,
			CurrentState: obj2,
		})

		_, _ = w.Write(str)
	}
}
