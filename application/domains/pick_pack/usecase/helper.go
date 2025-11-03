package usecase

const NO_RETRY = 0
const NO_DELAY = 0
const TIMEOUT_20S = 20

func apiRetryCondition(status int, body []byte, errRequest error) bool {
	return false
}
