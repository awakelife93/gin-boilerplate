package protocol

const (
	FORBIDDEN_STATUS             int = 403
	NOT_FOUND_STATUS             int = 404
	OK_STATUS                    int = 200
	CREATE_STATUS                int = 201
	BAD_REQUEST_STATUS           int = 400
	CONFLICT_STATUS              int = 409
	INTERNAL_SERVER_ERROR_STATUS int = 500
)

const (
	UNAUTHORIZED_MESSAGE          string = "UNAUTHORIZED" // 인증 실패
	FORBIDDEN_MESSAGE             string = "FORBIDDEN"
	NOT_FOUND_MESSAGE             string = "NOT_FOUND" // 리소스 및 데이터 없음
	OK_MESSAGE                    string = "OK"        // 정상 처리
	CREATE_MESSAGE                string = "CREATE"
	BAD_REQUEST_MESSAGE           string = "BAD_REQUEST"           // 잘못된 요청
	CONFLICT_MESSAGE              string = "CONFLICT"              // 리소스 및 데이터 중복
	INTERNAL_SERVER_ERROR_MESSAGE string = "INTERNAL_SERVER_ERROR" // 기타 예외 처리
)
