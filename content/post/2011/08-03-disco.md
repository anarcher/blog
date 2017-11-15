+++
date="2011-08-03"
tags=["disco","mapreduce"]
title = "Disco 사용기"
draft = false
+++

[Disco][Disco]을 잘 모르지만, 이번에 통계관련 계산에 한번 써봤다. 

## map/reduce

map/reduce는 map의 결과가 하나의 reduce에서 같은 key으로 모이는 것 보장한다.로직에 따라 다르겠지만.  전체 갯수을 얻는 것이 아니라면 reduce을 늘리는 것이 성능에 좋다. (당연한 이야기이지만)   

## combiner

Disco 문서에서는 보기 힘든데, Hadoop처럼 combiner을 만들 수 있다. combiner은 일종의 accumulator인데. 각 map마다 map의 결과를 받아서 reduce에 넘기기전에 buffer을 가진 function(combiner)을 만들 수 있는데. 대부분의 경우 combiner을 만드는 것이 속도나 메모리 사용에 좋을 것 같다.  

 
## ddfs chunk size 와 map 
ddfs에 데이터를 넣을때, 기본적으로 blob size은 64M이다. 문제는 mapper은 blob마다 할당되는데. 데이터가 좀더 넓게 펴져 있다면, 좀 더 많은 머신들이 map에 참여하게 된다. 
ddfs command에서는 chunk size을 바꿀수 없고. ddfs chunk으로 넣은 데이터를 원하는 size으로 짤라주는 작업이 필요하다. [^1]

## ddfs chunk 읽기 

ddfs chunk도 일종의 mapreduce job으로 돌린다. ddfs xcat command처럼 chunk의 데이터를 읽으려면 
disco_input_stream 으로 읽어야 한다. 
ddfs command도 python code이니. ddfs command code을 참고했다. 

## Disco vs Hadoop

대부분의 경우 단일 머신에서 계산을 비교하면 그냥 일반적인 계산 코드가 빠르다. 하지만 확장성이나 데이터를 분산해서 두고 규모가 커지더라도 대응할수 있게 만들려면 M/R이 나쁘지 않은 듯 하다. 

내가 하는 Jobs들은 대용량의 데이터라기 보다는 많은 Jobs이 참여하는 구조이라서 (더군다나 계층적인 구조.. OTL) job들을 chaining해서 이전 reduce의 output이 다음 map의 input이 되는 패턴이 계층적으로 되어지니까 코드가 귀찮아 진다. (핑계 없는 무덤을 없더라는 속담을 생각하며. -_-)


[Storm][Storm]이나 [Spark][Spark]에서는 좀더 이런 구조에 적절할지 궁금하다. 
그나저나 Storm은 이번달에 나오겠군. 


[Disco]: http://discoproject.com/
[Tutorial]: http://discoproject.org/doc/start/tutorial.html
[Spark]: https://github.com/mesos/spark
[Storm]: http://engineering.twitter.com/2011/08/storm-is-coming-more-details-and-plans.html

[^1]: Hadoop은 어떻게 작동하는지 잘 모르겠다. map의 partition을 줄수 있을 것 같은데. 처음에 당황했던 부분이다. 


