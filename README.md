# go-taskqueue-client

[go-taskqueue](https://github.com/<github계정>/go-taskqueue) 서버와 통신하는 독립 클라이언트입니다.

`go-taskqueue`의 `internal` 패키지를 import하지 않고, HTTP 통신만으로 서버와 상호작용합니다. 별도 모듈/별도 저장소로 분리했습니다.

## 하는 일

- `POST /jobs`로 작업을 등록하고, `GET /jobs/{id}`로 상태 변화(pending → processing → completed)를 조회합니다.
- `sync.WaitGroup` + goroutine을 이용해 다수의 작업을 동시에 등록하는 시나리오로, 서버의 워커 풀이 동시 요청 상황에서 어떻게 동작하는지 확인합니다.

## 실행 방법

먼저 [go-taskqueue](https://github.com/<github계정>/go-taskqueue) 서버를 `:8080`에서 실행한 뒤:

```bash
go run main.go
```

## 참고

이 클라이언트는 처리량/지연시간을 정밀하게 측정하는 도구가 아닙니다. 성능 측정은 `hey`, `k6` 같은 전문 부하 테스트 도구를 별도 환경에서 사용해야 합니다. 자세한 내용은 서버 저장소의 README를 참고하세요.
