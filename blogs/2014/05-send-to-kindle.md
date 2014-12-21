<!-- Title: Send-to-kindle 2014-05 -->
<!-- Tags: send-to-kindle,golang,concurrency,coding,http -->


> 읽어야 할 글들을 Kindle에 던지는데,뭘 킨들에 던졌는지 기록해 두면 좋을 것 같아서 적어본다. 

# Go's power is in emergent behavior

- http://www.onebigfluke.com/2014/04/gos-power-is-in-emergent-behavior.html

golang의 interface의 특징을 두가지 예(http.HandlerFunc,opaque type)을 가지고 설명. 
저번에 있었던 gophercon에서 들은 내용에 대한 글.


# Go Parallel 1,2 

- https://software.intel.com/en-us/blogs/2013/06/18/go-parallel
- https://software.intel.com/en-us/blogs/2014/04/13/go-parallel-2

고루틴에 대한 예제와 함께 짧게 고루틴을 알수 있는 글. 

# Go: Best Practices for Production Environments

- http://peter.bourgon.org/go-in-production/

soundcloud에서 go 프로젝트를 하면서 생긴 팁들을 나열.

개인적으로는 python의 virtualenv처럼 여러 GOPATH을 관리 할 수 있으면 좋겠다. 
(아마 이미 관련 도구가 있을 듯 하기도 하다.)

# Why I like Go?

- https://gist.github.com/freeformz/4746274
- https://news.ycombinator.com/item?id=5195257

# The Case for React.js and ClojureScript
 
- http://murilopereira.com/the-case-for-reactjs-and-clojurescript/#/

react.js와 clojurescript에 대한 이야기. 
react.js은 거칠게 이야기해서 MVC중 V에 해당한다. 
개인적으로 관심을 가지고 있는 프론트엔드 프로젝트중 하나. 


# Load balancing with Vulcan and Etcd

- http://vulcan.ghost.io/intro/

docker을 만든 dotcloud의 https://github.com/dotcloud/hipache는  redis을 backend data로 사용하는데 반해,

vulcan은 etcd에 있는 데이터를 기반으로 동적으로 로드밸랜싱을 해주는 리버스 프록시 프로젝트이다. 

go로 만든 프로젝트.


# Why we don't hire programmers based on puzzles, API quizzes, math riddles, or other parlor tricks

- http://signalvnoise.com/posts/3071-why-we-dont-hire-programmers-based-on-puzzles-api-quizzes-math-riddles-or-other-parlor-tricks

- http://gettingreal.37signals.com/ch08_Kick_the_Tires.php

실제 코드와 이슈로 좋은 프로그래머를 찾는 것이 특정 퀴즈로 찾는 것보다 성공률이 높다.

함께 일해보는 것 만큼 좋은 인터뷰는 없다.


# Immutable Infrastructure with Ansible and Packer

- http://blog.codeship.io/2014/05/08/ansible-packer-immutable-infrastructure.html


