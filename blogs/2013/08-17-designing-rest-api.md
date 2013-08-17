<!-- Title:Designing REST API -->
<!-- Tags: restapi -->

REST API는 RPC(Remote Procedure Call) API와 다르게 기본적으로 절차적인 명령 호출의 목록으로 설계하는 것이 아니라고 생각된다. 

REST(Representational State Transfer) API는 선언적인 방법으로 리소스(Resource)의 표현되는 상태를 변화(GET,POST,PUT,PATCH,DELETE...)시키는 인터페이스라고 할수 있다. 

그리고 이러한 리소스의 상태에 따라 시스템이 동작하는 구조라고 생각한다.(직접 시스템의 동작을 명령하는 인터페이스를 가진 것이 아닌.) 

REST API을 설계할때, RPC API 설계개념을 가져와 작성하게 되면 HTTP Methods(GET,POST,PUT...)으로 동작으로 명령을 정의하게 되어서 표현제약이 발생하고,결국 리소스을 가져다가 특정 동작을 정의하게 된다. 

책 "일관성 있는 웹 서비스 인터페이스 설계를 위한 REST API 디자인 규칙"에서는 리소스 프로토타입에서 컨트롤러라는 이름으로 이러한 절차적인 호출을 위한 프로토타입이 있기는 하다. 

하지만 기본적으로 리소스의 상태를 변화시키는 방법으로는 만들수 없는 경우에 사용되는 프로토타입이라고 생각된다.


