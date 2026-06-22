package response

import (
	"net/http"
	"testing"
)

func TestSuccess(t *testing.T) {
	data := map[string]string{"key": "value"}
	resp := Success(data)

	if resp.Code != http.StatusOK {
		t.Errorf("Success() Code = %d, want %d", resp.Code, http.StatusOK)
	}
	if resp.Message != "success" {
		t.Errorf("Success() Message = %q, want %q", resp.Message, "success")
	}
	if !resp.Success {
		t.Error("Success() Success should be true")
	}
	if resp.Data == nil {
		t.Error("Success() Data should not be nil")
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		code    int
		message string
	}{
		{400, "参数错误"},
		{401, "未授权"},
		{403, "禁止访问"},
		{500, "服务器错误"},
	}

	for _, tt := range tests {
		resp := Error(tt.code, tt.message)
		if resp.Code != tt.code {
			t.Errorf("Error(%d, %q) Code = %d, want %d", tt.code, tt.message, resp.Code, tt.code)
		}
		if resp.Message != tt.message {
			t.Errorf("Error(%d, %q) Message = %q, want %q", tt.code, tt.message, resp.Message, tt.message)
		}
		if resp.Success {
			t.Error("Error() Success should be false")
		}
	}
}

func TestPage(t *testing.T) {
	list := []string{"a", "b", "c"}
	resp := Page(list, 1, 10, 30)

	if resp.Code != http.StatusOK {
		t.Errorf("Page() Code = %d, want %d", resp.Code, http.StatusOK)
	}
	if !resp.Success {
		t.Error("Page() Success should be true")
	}

	pageData, ok := resp.Data.(*PageResponse)
	if !ok {
		t.Fatal("Page() Data should be *PageResponse")
	}

	if pageData.Page != 1 {
		t.Errorf("Page() Page = %d, want 1", pageData.Page)
	}
	if pageData.PageSize != 10 {
		t.Errorf("Page() PageSize = %d, want 10", pageData.PageSize)
	}
	if pageData.Total != 30 {
		t.Errorf("Page() Total = %d, want 30", pageData.Total)
	}
	if pageData.TotalPage != 3 {
		t.Errorf("Page() TotalPage = %d, want 3", pageData.TotalPage)
	}
}

func TestToRPCResponse(t *testing.T) {
	resp := &Response{
		Code:    200,
		Message: "success",
		Success: true,
	}

	rpcResp := resp.ToRPCResponse()
	if rpcResp.Code != 200 {
		t.Errorf("ToRPCResponse() Code = %d, want 200", rpcResp.Code)
	}
	if rpcResp.Message != "success" {
		t.Errorf("ToRPCResponse() Message = %q, want %q", rpcResp.Message, "success")
	}
	if !rpcResp.Success {
		t.Error("ToRPCResponse() Success should be true")
	}
}

func TestNewRPCResponse(t *testing.T) {
	resp := NewRPCResponse(200, "ok", true)
	if resp.Code != 200 {
		t.Errorf("NewRPCResponse() Code = %d, want 200", resp.Code)
	}
	if resp.Message != "ok" {
		t.Errorf("NewRPCResponse() Message = %q, want %q", resp.Message, "ok")
	}
	if !resp.Success {
		t.Error("NewRPCResponse() Success should be true")
	}
}
