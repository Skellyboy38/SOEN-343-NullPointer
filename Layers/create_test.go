package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/handler"
	"github.com/Skellyboy38/SOEN-343-NullPointer/Layers/domain_layer/mappers"
)

func BenchmarkReservation(b *testing.B) {

	mappers.Init()
	mappers.InitUOW()

    for i := 0 ; i < b.N ; i++{
        data := url.Values{}
        data.Set("dataRoom", "1")
        data.Add("userID", "1111111")
        data.Add("startTime", "2016-12-22 12:00:00")
        data.Add("endTime", "2016-12-22 16:00:00")

        req, err := http.NewRequest(
            http.MethodPost,
            "http://localhost:9000/createReservation",
            bytes.NewBufferString(data.Encode()),
        )
        req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

        if err != nil {
            b.Fatalf("could not create request %v", err)
        }

        rw := httptest.NewRecorder()

        handler.CreateReservation(rw, req)

        if !strings.Contains(rw.Body.String(), "Success") {
            if rw.Code != http.StatusOK {
                b.Errorf("expected status %s; got %d", http.StatusOK, rw.Code)
            }
        }
        rw.Flush()

        data = url.Values{}
        data.Set("dataRoom", "1")
        data.Add("userID", "1111111")

        reqByUser, err := http.NewRequest(
            http.MethodPost,
            "http://localhost:9000/reservationsByUser",
            bytes.NewBufferString(data.Encode()),
        )

        reqByUser.Header.Add("Content-Type", "application/x-www-form-urlencoded")

        if err != nil {
            b.Fatalf("could not create request %v", err)
        }

        userRecorder := httptest.NewRecorder()

        handler.GetReservationsByUserID(userRecorder, reqByUser)

        if !strings.Contains(userRecorder.Body.String(), "2016-12-22T12:00:00Z") {
            b.Errorf("expected status %s; got %s", "2016-12-22 12:00:00", userRecorder.Body.String())
            if userRecorder.Code != http.StatusOK {
                b.Errorf("expected status %s; got %d", http.StatusOK, userRecorder.Code)
            }
        }
    }
}
