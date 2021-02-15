package gochatwork

import (
	"encoding/json"
	mock_moment "github.com/go-chatwork/mock"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

const ApiKey = ``

func TestMe(t *testing.T) {
	var client *Client

	t.Run("実体を利用したテスト", func(t *testing.T) {
		client = NewClient(ApiKey)
		me := client.Me()
		if me.AccountId == 0 {
			t.Error("実体を使ったテストで失敗した。")
		}
		// テスト出力
		// fmt.Printf("%+v\n", me)
	})

	t.Run("モックを利用したテスト", func(t *testing.T) {
		client = NewClient(ApiKey)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		mockHttpClient := mock_moment.NewMockHttp(ctrl)

		// me のリクエストのレスポンス例
		meResponse := []byte(`{"account_id":123,"room_id":12345,"name":"cw-bot","chatwork_id":"","organization_id":4321,"organization_name":"cw","department":"","title":"","url":"","introduction":"bot desu","mail":"","tel_organization":"","tel_extension":"","tel_mobile":"","skype":"","facebook":"","twitter":"","avatar_image_url":"","login_mail":"abcd@test.com"}`)

		// me のリクエストを受け取ったら レスポンス例のデータを返す。
		mockHttpClient.
			EXPECT().
			Get("/me", map[string]string{}).
			Return(meResponse).
			MinTimes(1).
			MaxTimes(1)
		// チャットワーククライアント内の HTTP クライアントをモックに差し替える。
		client.InnerHttpClient = mockHttpClient

		data := client.Me()

		// テスト結果比較用にレスポンス例をパースする。
		var expectedMe Me
		json.Unmarshal(meResponse, &expectedMe)

		if ! reflect.DeepEqual(data, expectedMe) {
			t.Error("モックを使ったテストで失敗した。")
		}
		if data.AccountId != 123 {
			t.Error("期待したアカウントIDとちがう。")
		}
		// テスト出力
		// fmt.Printf("%+v\n", data)
	})
}
