<!-- Title:Docker deployment slide -->
<!-- Tags: docker,slide -->

사내에서 Docker 사용에 대한 이야기가 나와서 간단하게 조사해보았다. 
매력적인 프로젝트인것은 확실하지만, 기존에 chef나 opsworks로 구성되어 있는 부분도 있고 현재 시스템 구조상 
docker으로 이전해서의 이득이 생각외로 크지 않다는 것이 중평.

docker의 장점은 대개 두가지로 이야기 할수 있겠다.

- Portability (이식성)
- Repeatability (반복성)

이식성(Portability)은 가장 크게 얻는 이점은  개발환경과 실제 프로덕션 환경의 차이를 설정변수 몇개로 축소할수 있다는 점이라고 생각한다. 
여러 환경(개발,테스트,프로덕션)을 일치시키기 위한 별다른 노력을 할 필요 없다는 점이 가장 큰 매력으로 생각한다. 
물론 자바와 비슷하게 한번 작성하면 어디서든 동작한다는 것이 여러 Host 환경에서 동작할 수 있다는 점도 무시할수는 없겠다.

반복성(Repeatability)은 결국 컨테이너와 이미지를 생성/변경/실행이 빠르고 가볍기 때문에 지속적인 통합(Continuous Integration)과 배포(Deployment) 이득이 있다. 

아래 장표에서 밴치마크 부분은 docker 0.7/linux 3.8 기준으로 작성되어 있는 것을 참고했다. linux 3.11에서 network부분의 향상이 있어서 지금 docker 1.0/linux 3.11 에서의 network 성능은 Host와의 차이가 크지 않을 것이라고 생각한다. (자료를 찾아보거나 테스트해 볼 필요는 있어보인다)

<iframe src="//slides.com/anarcher/docker-deployment/embed" width="576" height="420" scrolling="no" frameborder="0" webkitallowfullscreen mozallowfullscreen allowfullscreen></iframe>


