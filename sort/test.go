package main

import (
	"fmt"
	"time"
	"sort"

)

type Pod struct {
  podName string
  creationTime time.Time
}

func main() {

	var pod1 Pod
	var pod2 Pod
	var pod3 Pod
	var pod4 Pod
	
  

	pod1.podName = "mypod2"
	pod1.creationTime = time.Now()

	time.Sleep(2 * time.Second)


	pod2.podName = "mypod2"
	pod2.creationTime = time.Now()

	time.Sleep(2 * time.Second)
  

	pod3.podName = "mypod2"
	pod3.creationTime = time.Now()

	time.Sleep(2 * time.Second)


	pod4.podName = "mypod2"
	pod4.creationTime = time.Now()

	podList := []Pod {pod2,pod4,pod1,pod3}

	for _, element := range podList {
		fmt.Println(element.podName, "--", element.creationTime)
	}

	fmt.Println("====================================")

	filterPod := []Pod {};

	for _,element := range podList {
		if(element.podName == "mypod2") {
			 filterPod = append(filterPod,element);
		}
	}

	
	for _, element := range filterPod {
		fmt.Println(element.podName, "--", element.creationTime)
	}

	fmt.Println("====================================")

	sort.Slice(filterPod, func(i, j int) bool {
		return filterPod[i].creationTime.Unix() > filterPod[j].creationTime.Unix()
	})

	for _, element := range filterPod {
		fmt.Println(element.podName, "--", element.creationTime)
	}


	fmt.Printf("latest pod time : %s",filterPod[0].creationTime)




}
