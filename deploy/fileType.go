package deploy

type FileModel struct {
}

const (
	Deployment string = `---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: [name]
  namespace: [namespace]
spec:
  replicas: 1
  template:
    metadata:
      namespace: [namespace]
      labels:
        app: [name]
    spec:
      containers:
      - name: [name]
        image: [url]/[author]/[name]:v[version]
        imagePullPolicy: Always
        resources:
          # keep request = limit to keep this container in guaranteed class
          limits:
            cpu: [cpuLimit]
            memory: [memoryLimit]
          requests:
            cpu: [cpuRequest]
            memory: [memoryRequest]
        volumeMounts:
        - mountPath: [logPath]
          name: [name]
        command: ["/bin/sh", "-c"]
        args: ["[cmdArgs]"]
      volumes:
      - hostPath:
          path: [logTargetPath]
        name: [name]
`
//	Deployment string = `---
//apiVersion: v1
//kind: ReplicationController
//metadata:
//  name: [name]
//  namespace: [namespace]
//  labels:
//    app: [name]
//    version: v1.0
//spec:
//  replicas: 1
//  selector:
//    app: [name]
//    version: v1.0
//  template:
//    metadata:
//      namespace: [namespace]
//      labels:
//        app: [name]
//        version: v1.0
//    spec:
//      containers:
//        - name: [name]
//          image: [url]/[author]/[name]:v[version]
//          imagePullPolicy: Always
//          resources:
//            limits:
//              cpu: [cpuLimit]
//              memory: [memoryLimit]
//            requests:
//              cpu: [cpuRequest]
//              memory: [memoryRequest]
//          volumeMounts:
//            - name: [name]
//              mountPath: [logPath]
//          command: ["/bin/sh","-c"]
//          args: ["[cmdArgs]"]
//      volumes:
//        - name: [name]
//          hostPath:
//            path: [logPath]
//`

	Ingress string = `---
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name:  [name]
  namespace: [namespace]
spec:
  rules:
  - host:  [domain]
    http:
      paths:
      - path: /
        backend:
          serviceName: [name]
          servicePort: [port]
`

	Service string = `---
apiVersion: v1
kind: Service
metadata:
  name: [name]
  namespace: [namespace]
  [annotations]
  labels:
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: "[name]"
spec:
  type: NodePort
  ports:
    - name: [name]
      port: [port]
      targetPort: [servicePort]
      protocol: TCP
  selector:
    app: [name]
`

	ServiceDev string = `---
apiVersion: v1
kind: Service
metadata:
  name: [name]
  namespace: [namespace]
  labels:
    kubernetes.io/cluster-service: "true"
    kubernetes.io/name: "[name]"
spec:
  type: NodePort
  ports:
    - name: [name]
      port: [servicePort]
      targetPort: [port]
      protocol: TCP
      nodePort: [exportPort]
  selector:
    app: [name]
`

	Dockerfile string = `FROM reg.miz.so/library/centos:7
MAINTAINER "Eno <eno@hyx.com>"

#修改容器时区
RUN rm /etc/localtime && \
ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /[name]/build

ADD ./build/main /[name]/build/[name]
ADD ./build/config.json /[name]/build/config.json

RUN mkdir /data && cd /data && mkdir logs

CMD ["./[name]", "-conf", "config"]
`
)
