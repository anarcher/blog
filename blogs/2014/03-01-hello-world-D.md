<!-- Title: D 언어 사용기 -->
<!-- Tags: dlang -->
<!-- Created: 2014-03-01 -->
<!-- Updated: 2014-03-01 -->

아직 서비스에 사용이 가능할지는 미지수이지만, Python으로 구성한 REST API Server을 이번에 D언어로 프로토타이핑을 하게 되었다. 

우선 사용한 라이브러리는 [vibe.d](http://vibed.org/)와 Backend Server와의 통신을 위해 [thrift](http://thrift.apache.org/)이다. 

[vibe.d](http://vibed.org)은 기본적으로 [redis](http://redis.io)와 [mongodb](http://www.mongodb.org/)을 지원하지만, Thrift에 대한 지원은 없다. 

vibed의 [Writing native database drivers - and a new MySQL driver port](http://vibed.org/blog/posts/writing-native-db-drivers) 을 참고하여 TSocket을 vibe.d의 TCPConnection을 사용하도록 만들어서 사용했다. 

(지금으로써는) 간단한 REST API Server이라서 언어적인 기능을 많이 사용하지는 않았지만,간단히 D언어 사용에 대한 평을 해보자면,

- 타입시스템에 반하여 암묵적으로 컴파일러가 처리하는 부분이 없다. 
- 함수적 프로그래밍(Functional Programming)위해 만들지 않고, 함수적 프로그래밍의 개념을 절차적인 언어의 기반에서 적절하게 적용했다. 
- 물론 객체지향 프로그래밍을 지원하며, 객체지향적인 프로그래밍에 필요한 참조기반의 의미론과 이에 반해 struct등을 사용한 값기반의 의미론을 모두 지원한다. 
- 가비지콜렉션을 지원하지만, 명시적으로 메모리 할당과 해제를 할수도 있다. (하지만 난 거의 그렇게 사용하지 않았다)
- 동시성에 대한 설계에서 erlang등의 ActorModel에 영향을 받아서 설계되었다. (하지만 다른 언어처럼 잠금기반의 라이브러리도 지원한다) 

특히 [The D Programming language](http://www.amazon.com/The-Programming-Language-Andrei-Alexandrescu/dp/0321635361) 이 책은 D 언어에 관심이 없더라도, 한번쯤 읽어 볼만 하다.저자가 언어의 디자인 결정을 왜 했는지에 대한 이야기를 정말 재미있게 풀었다.

