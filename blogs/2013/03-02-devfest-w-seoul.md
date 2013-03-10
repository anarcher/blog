<!-- Title:DevFest W Seoul -->
<!-- Tags: devfest -->
<!-- Created: 2013-03-02 -->
<!-- Updated: 2013-03-02 -->


오랜만에 데브 페스트에 놀려 갔다.
https://sites.google.com/site/2013devfestwkorea/home/aboutspeaker


## 구글인앱빌링 사용 모듈 구현사례

안드로이드 프로그래밍을 해본 경험이 없어서,잘 알지 못하는 분야이다. 
구글의 인앱빌딩 지원을 버젼별로 상세히 설명되어있고,적어도 (나처럼 무경험자에게는)
어떻게 동작하는지 알수 있는 자리이었다. 

나중에 안드로이드 프로그래밍을 한다면,그때 다시 한번 찾아 볼듯 하다. 

## Go Lang for Java Programmer

GoLang에 대한 기본적인 문법과 개념을 설명했다. 
발표에 대한 연습을 많이 했다는 걸 느낄수 있었다. 

사실 나는 Go에 관심이 있어서 DevFest에 참가한 것이라고 할수 있는데. 기본적인 GoLang에 대한 차분한 설명을 잘 들을수 있었다. 

## Chrome UI Frameworks

크롬OS의 Arua UI에 대한 기본적인 설명을 하는 세션이었다. 
Arua UI는 원도우 UI 시스템인데, 사실 난 크롬OS에서 크게 감흥이 없어서 그런지 UI을 구경하는 자리이었던것 같다. 

기억이 남는 부분은 GtkWindow을 쓰지않고 구글은 독자적으로 Arua UI을 만들었다는 점. 
기능한 어떤 이점이 있는지 모르겠다. 

## How to live like goer

가장 재미있게 들은 발표이다. 
Python에서 Go으로 넘어가면서의 기록이라고도 할수 있는데. 무언가 정확한 정보를 전달하는 것보다 
사변적으로 이런 저런 다른 이의 생각과 경험을 알수 있어서 그런지 가장 즐거웠다. 


## AnguarJS

사실 Web Frontend에서 ServerSideRendering와 ClientSideRendering에 대해 갑을논박이 많은 듯 하다.
이부분은 사실 IRC(오징어의 #codeport에 놀려 오시면 종종 이런 이야기하는 무리?를 발견할수 있다.)에서 사람들과 이야기 하곤 했는데. 
ClientSideRendering에서 Backbone.js보다 더 요즘 인기를 얻고 있는 프레임워크이라서 관심이 있었다. 

Backbone.js보다 좀더 일체화된 구성을 가지고 있는데. 많은 부분이 구성이 되어 있어서 Backbone.js보다 좀더 쉽게 시작할수 있는 것 같다. 

요즘 Backbone.js에서는 view dispose에 대한 이슈로 패치가 추가되었는데. AnguarJS에서는 어떤식으로 객체의 확득과 소멸을 하는지 궁금해졌다. 

특히 One page application에서는 JS의 메모리 관리가 큰 이슈가 된다고 생각한다. 


## Concurrency is not parallelism

RobPike의 "Concurrency is not parallelism"을 GoLangUserGroup에서 다시 발표한 것이다. 
CSP의 개념을 가지고 만든 동시성을 지원하는 Golang의 goroutine으로 병렬성을 어떻게 얻는지에 대한 이야기이다. 

Concurrency은 프로그램의 특성이고,Parallel은 기계의 성질이다. 어떤 프로그램이 동시에 여러 루틴이 돌아간다면 Concurrency한 것이고,이 프로그램이 정말 동시에 동작할지는 사실 머신에 따라 다르다. 
싱글 코어에서도 Concurrent하게 돌겠지만(시분할), 정말 병렬적으로 돌려면, 멀티코어상에서 가능하다. 

GoLang의 동시성 지원은 멀티코어 시대의 위기?를 해결하고자 이슈를 얻고 있는 함수형 프로그래밍(물론 이것 하나때문에 요즘 관심을 받는 것은 아니지만)에 대한 생각을 다시 하게 만들었다. 

어떤 goroutine들은  Concurrent을 한 Thread에서만 획득하고 싶다거나(ThreadContention을 피하고 싶을 경우라든가) 하는 구성을 어떻게 하는지 궁금해졌다. 
예를 들어,goroutine으로 PythonGenerator을 구성한다면 말이다. 

## 오랜만의 컨퍼런스 

사실 3일연휴이라서 여유가 있기에 올수 있었다. 나도 컨퍼런스에 대한 맘이 예전같지 않다.(응?;)
나름 지인들도 보고 공짜 점심도 먹을 수 있어서..;ㅁ; 즐거운 하루이었던것 같다. 
아래는 발표자료 링크이다. 

https://docs.google.com/a/enswer.net/folder/d/0B82z67M0_dLHWXdfNEtMQ2hYSXM/edit?usp=sharing



