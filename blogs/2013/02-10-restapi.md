<!-- Title:REST API 디자인 룰 북 -->
<!-- Tags: restapi -->
<!-- Created: 2013-02-10 -->
<!-- Updated: 2013-02-10 -->

HTTP 위에 JSON 기반으로 API을 만드는 일이 근래 많아졌다. 
REST API가 최근에 이야기되는 개념도 아니지만, API을 만들때 생각외로 고민해야 할 부분들이 적지 않다. 

우선 HTTP Methods (GET,POST,PUT,HEAD,PATCH,DELETE)만 가지고 Resource의 행동을 다 표현하기에는 부족함이 있는 것이 아닌가 싶기도 하다. 

또 Client에서 GET과 POST을 지원할 경우(혹은 다른 메소드를 지원하는 불투명한 경우)에 쉽게 HTTP Methods을 결정하기 어려워지기도 하다. [^1]

사실 리소스 모델링 부분이 어려운데[^2], [RuleBook][RuleBook]을 참고하여 구성하고 있다.

[RuleBook][RuleBook]에서 인상적인 부분들은 다음과 같다. 

- Resource의 프로토타입을 제시 
   - 도큐먼트 
   - 컬렉션 
   - 스토어 
   - 컨트롤러 




[RuleBook]: http://shop.oreilly.com/product/0636920021575.do

[^1]: Rails의 경우 (몇년 전 기억이지만) 지원하지 않은 브라우저를 위해  form_for 태그에 _method으로 추가해서 사용할수 있도록 했다. (http://apidock.com/rails/ActionView/Helpers/FormHelper/form_for)

[^2]: 사실 XML-RPC가 REST API보다 강제적이라서 좀더 클라이언트의 검증등의 이점이 있지 않나 하는 생각도 있다. 



