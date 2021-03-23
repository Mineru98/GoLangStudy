package main

import "os"

/*
	panic은 일반적으로 무언가가 예상치 못하게 잘못되었음을 의마한다.
	대부분 정상적인 작동 중에는 발생해선 안되는 오류 또는 정상적으로 처리할 준비가 되어있지
	않은 오류에 대해 빠르게 실패시키는데 사용한다.
*/

func main() {
	/*
		예상치 못한 에러를 체크하기 위해 사이트 전체에서 패닉을 사용한다.
		이는 사이트에서 패닉을 하기 위해 설계된 유일한 프로그램이다.
	*/
	panic("a problem")

	/*
		패닉의 일반적인 사용은 어떤 함수가 어떻게 처리할지 모르는(또는 하고 싶지 않은)
		에러값을 반환했을때 중단을 하기 위함이다.
		다음은 파일을 생성할 때 예상치 못한 에러가 발생할 경우 패니킹을 하는 예시이다.
	*/
	_, err := os.Create("./file")
	if err != nil {
		panic(err)
	}
}

/*
	프로그램을 실행하면 패닉을 일으키며, 에러 메시지와 고루틴 트레이스 정보,
	그리고 0이 아닌 상태값과 함계 종료된다.

	참고로 에러를 핸들링 하기 위해 인셉션을 사용하는 타 언어와는 달리,
	Go에서는 에러 가능성이 있는 곳에서 에러를 가리키는 반환값을 사용하는게 일반적이다.
*/
