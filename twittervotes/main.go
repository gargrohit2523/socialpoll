package main

func main() {
	// read data from mongo
	// pass data to twitter read module and read tweets
	// publish the tweets to nsq
	// then a goroutine shall pick message from nsq
	// send the message to mongo
}
