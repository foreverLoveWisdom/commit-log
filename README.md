### Introduction

- Go has become the most popular language for building distributed services, as shown by projects like Docker, Etcd, Vault, CockroachDB, Prometheus, and Kubernetes.
- Despite the number of prominent projects such as these, however, there’s no resource that teaches you why or how you can extend these projects or build your own.
- Where do you begin if you want to build a distributed service? When I began learning how to build distributed services, I found the existing resources to be of two extremes: • Concrete code—distributed services are large, complex projects, and the prominent ones have had teams working on them for years.

- The layout of these projects, their technical debt, and their spaghetti code bury the ideas you’re interested in, which means you have to dig them out.
- At best, learning from code is inefficient.
- Plus there’s the risk that you may uncover outdated and irrelevant techniques that you’re better off avoiding in your own projects.

- Abstract papers and books—papers and books like Designing Data- Intensive Applications by Martin Kleppmann1 describe how the data structures and algorithms behind distributed services work but cover them as discrete ideas, which means you’re left on your own to connect them before you can apply them in a project.
- These two extremes leave a chasm for you to cross.
- I wanted a resource that held my hand and taught me how to build a distributed service—a resource that explained the big ideas behind distributed services and then showed me how to make something of them.
- I wrote this book to be that resource. Read this book, and you’ll be able to build your own distributed services and contribute to existing ones.

#### Chapter 11: Deploy Applications with Kubernetes to the Cloud

- Congratulations! You deployed your service to the cloud. Now any person on the Internet can use your service. You set up a Google Cloud account, a project, and a GKE cluster. You also learned how to write a simple controller to extend the behavior of Kubernetes resources with Metacontroller. We’ve now reached the end of the book, and you’ve accomplished a lot. You’ve made a distributed service from scratch. You’ve learned distributed computing ideas like service discovery, consensus, and load balancing. You’re ready to make your own distributed services and contribute to existing projects. Go leave your mark on this growing field!

#### Chapter 10: Deploy Applications with Kubernetes Locally

- In this chapter, you learned the fundamentals of Kubernetes and how to use Kind to set up a Kubernetes cluster that you can run on your machine or on a CI. You also learned how to create a Helm chart and how to install your Helm chart into Kubernetes to run a cluster of your service. You learned quite a lot! In the next chapter, we’ll build on this knowledge and deploy your service on a cloud platform.

#### Chapter 9: Discover Servers and Load Balance from the Client

- Now you know how gRPC resolves services and balances calls across them and how you can build your own resolvers and pickers. You can write your own resolver so that your clients dynamically discover servers. And you saw how pickers are useful for more than just load balancing—you can build your own routing logic with them. In the next part of the book, we’ll look at how to deploy our service and make it live.

#### Chapter 8: Coordinate Your Services with Consensus

- In this chapter, you learned how to coordinate distributed services with Raft by adding leader election and replication to our service. We also looked at how to multiplex connections and run multiple services on one port. Next, we’ll talk about client-side discovery, so clients can discover and call our servers.

#### Chapter 7: Server-to-Server Service Discovery

- Now when our servers discover other servers, they replicate each other’s data. That’s a problem with our replication implementation: when one server dis- covers another, they replicate each other in a cycle! You can verify it by adding this code at the bottom of your test:

  ```go
  consumeResponse, err = leaderClient.Consume(
  context.Background(),
  &api.ConsumeRequest{
  Offset: produceResponse.Offset + 1,
  },
  )
  require.Nil(t, consumeResponse)
  require.Error(t, err)
  got := grpc.Code(err)
  want := grpc.Code(api.ErrOffsetOutOfRange{}.GRPCStatus().Err())
  require.Equal(t, got, want)
  ```

  - We only produced one record to our service, and yet we’re able to consume multiple records from the original server because it’s replicated data from another server that replicated its data from the original server. No, Leo, we do not need to go deeper.

#### Chapter 6: Observe Your Systems

- In this chapter, you learned about observability and its role in making reliable systems. You’ll find tracing especially useful in distributed systems, as it gives you a complete story of requests that take part over multiple services. You also learned how to make your service observable. Next, we’ll make our server support clustering to the service highly available and scalable.

#### Chapter 5: Secure Your Services

- You’ve learned how to secure services in three parts: by encrypting connections with TLS, through mutual TLS authentication to verify the identities of clients and servers, and by using ACL-based authorization to permit client actions. Next we’ll make our service observable by adding metrics, logs, and traces.

#### Chapter 4: Serve Requests with gRPC

- You now know how to define a gRPC service in protobuf, compile your gRPC protobufs into code, build a gRPC server, and test that everything works end- to-end across your client and server. You can build a gRPC server and client, and you can use your log over the network. Next we’ll improve the security of our service by encrypting the data sent between the client and server with SSL/TLS, and authenticating requests so we can know who’s making each request and whether they’re allowed to.

#### Chapter 3: Write a Log Package

- You now know what logs are, why they’re important, and how they’re used in various applications including distributed services. And then you learned how to build one! This log serves as the foundation of our distributed log. Now we can build a service on our library and make the library’s functionality accessible to people on other computers.

#### Chapter 2: Structure Data with Protocol Buffers

- In this chapter, we covered the protobuf fundamentals we’ll use throughout our project. These concepts will be vital throughout our project, especially as we build our gRPC client and server. Now let’s create the next vital piece of our project: a commit log library.

#### Chapter 1: Let’s Go

- In this chapter, we built a simple JSON/HTTP commit log service that accepts and responds with JSON and stores the records in those requests to an in-memory log. Next, we’ll use protocol buffers to manage our API types, generate custom code, and prepare to write a service with gRPC—an open source, high- performance remote procedure call framework that’s great for building dis- tributed services.
