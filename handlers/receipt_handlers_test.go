package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/kiramishima/receipt-processor/domain"
	"github.com/kiramishima/receipt-processor/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/unrolled/render"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReceiptHandlers_ReceiptProcessHandler(t *testing.T) {
	testCases := map[string]struct {
		ID            any
		buildStubs    func(uc *mocks.MockIReceiptService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
		item          map[struct {
			TotalWords       int
			TotalCents       int
			Total25          int
			Total2Items      int
			TotalItemDesc    int
			TotalOddDay      int
			TotalBetweenTime int
			TotalPoints      int
		}]domain.ReceiptBase
	}{
		"OK": {
			ID: 1,
			buildStubs: func(uc *mocks.MockIReceiptService) {
				uc.EXPECT().
					StoreReceipt(gomock.Any(), gomock.Any()).
					Times(1).
					Return("", nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
			item: map[struct {
				TotalWords       int
				TotalCents       int
				Total25          int
				Total2Items      int
				TotalItemDesc    int
				TotalOddDay      int
				TotalBetweenTime int
				TotalPoints      int
			}]domain.ReceiptBase{{TotalWords: 6, TotalCents: 0, Total25: 0, Total2Items: 10, TotalItemDesc: 6, TotalOddDay: 6, TotalBetweenTime: 0, TotalPoints: 28}: {
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []*domain.ReceiptItemBase{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			}},
		},
		"Missing Retailer Value": {
			ID: 1,
			buildStubs: func(uc *mocks.MockIReceiptService) {
				uc.EXPECT().
					StoreReceipt(gomock.Any(), gomock.Any()).
					AnyTimes().
					Return("", errors.New("Field: Retailer, Error: required\n"))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
			item: map[struct {
				TotalWords       int
				TotalCents       int
				Total25          int
				Total2Items      int
				TotalItemDesc    int
				TotalOddDay      int
				TotalBetweenTime int
				TotalPoints      int
			}]domain.ReceiptBase{{TotalWords: 6, TotalCents: 0, Total25: 0, Total2Items: 10, TotalItemDesc: 6, TotalOddDay: 6, TotalBetweenTime: 0, TotalPoints: 28}: {
				Retailer:     "",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []*domain.ReceiptItemBase{
					{
						ShortDescription: "Mountain Dew 12PK",
						Price:            "6.49",
					}, {
						ShortDescription: "Emils Cheese Pizza",
						Price:            "12.25",
					}, {
						ShortDescription: "Knorr Creamy Chicken",
						Price:            "1.26",
					}, {
						ShortDescription: "Doritos Nacho Cheese",
						Price:            "3.35",
					}, {
						ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
						Price:            "12.00",
					},
				},
				Total: "35.35",
			}},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mocks.NewMockIReceiptService(ctrl)
			tc.buildStubs(uc)

			recorder := httptest.NewRecorder()

			url := "/receipts/process"
			var item domain.ReceiptBase
			for _, obj := range tc.item {
				item = obj
			}
			// marshall data to json (like json_encode)
			marshalled, err := json.Marshal(item)
			if err != nil {
				log.Fatalf("impossible to marshall form: %s", err)
			}

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(marshalled))
			assert.NoError(t, err)

			router := chi.NewRouter()
			logger, _ := zap.NewProduction()
			slogger := logger.Sugar()
			r := render.New()
			NewReceiptHandlers(router, slogger, uc, r)
			router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}

func TestReceiptHandlers_ReceiptGetPointsHandler(t *testing.T) {
	var uids = []string{uuid.New().String(), uuid.New().String()}

	testCases := map[string]struct {
		ID            string
		buildStubs    func(uc *mocks.MockIReceiptService)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		"OK": {
			ID: uids[0],
			buildStubs: func(uc *mocks.MockIReceiptService) {
				uc.EXPECT().
					RetrieveReceipt(gomock.Eq(uids[0])).
					Times(1).
					Return(&domain.Result{
						ID:     uids[0],
						Points: 28,
					}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		"Not Found": {
			ID: uids[1],
			buildStubs: func(uc *mocks.MockIReceiptService) {
				uc.EXPECT().
					RetrieveReceipt(gomock.Any()).
					AnyTimes().
					Return(nil, errors.New(fmt.Sprintf("element with id: %s don't found", uids[1])))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mocks.NewMockIReceiptService(ctrl)
			tc.buildStubs(uc)

			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/receipts/%s/points", tc.ID)

			request, err := http.NewRequest(http.MethodGet, url, nil)
			assert.NoError(t, err)

			router := chi.NewRouter()
			logger, _ := zap.NewProduction()
			slogger := logger.Sugar()
			r := render.New()
			NewReceiptHandlers(router, slogger, uc, r)
			router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}

}
