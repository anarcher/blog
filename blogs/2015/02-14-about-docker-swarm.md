
<!-- Title:About Docker Swarm -->
<!-- Tags: docker,presentation -->

오랜만에 `docker-korea` [meetup](http://onoffmix.com/event/40731)에서 [About Docker Swarm](speakerdeck.com/anarcher/about-docker-swarm)이라는 이름으로 작은 발표를 했다. 

장표는 매우 간단하게 구성하고 데모에서 많은 이야기를 하고 싶었는데. 프로젝터설정부터 힘들어서 준비했던것을 절반밖에 이야기하지 못한것 같다. 

`Docker Swarm`은 다른 클러스텅 시스템과 다른 부분들이 있다. 가장 큰 다른 점은 `Docker`에서 만들었다는 것. 즉 `coreos/fleet`이 `Distributed Init System`을 표방하는 것과 비슷하게 `Swarm`은 `Docker-native Clustering System`을 이야기한다. [^1]

즉 다음과 같은 확장으로 `Docker` Interface(CLI/API)을 유지하고 싶어 한다. 

    docker run -d nginx-logger -e affinity:container==nginx nginx-logger

Google의 `k8s`나 다른 클러스트링 도구는 각자 자신만의 개념과 그 개념의 이해를 사용자에게 수반하게 만드는데. `Swarm`의 가장 큰 장점은 이러한 학습을 최소화 할수 있다는 점이 아닐까 싶기도 하다. [^2]

물론 이러한 접근이 쉽게 풀수 있는 문제를 어렵게 풀수 있는 여지가 있기도 하다. 

개인적으로는 `Docker`의 성공이 기술이 아니라 일종의 사용성에 기인한거라면 `Swarm`도 이러한 장점을 잘 이해하고 있다는 생각도 든다. 

ps) 하지만 아직 `0.1.0 (HEAD)`이고 프로덕션에서 쓰기에는 어리다. 

[^1]: 우리가 무언가 만들때 발생할수 밖에 없는 개념들이 실제 문제의 해결이 아니라 또 다른 문제를 야기할수도 있다는 생각도 들고, `coreos`은 `Unix`가 가지고 있는 개념을 확장하고 싶어하는 느낌이다. 
[^2]: 물론 `k8s`의 Pod라는 개념은 나쁘지 않다. 

