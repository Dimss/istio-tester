FROM centos:7
RUN yum -y update
RUN yum -y install bind-utils telnet vim
RUN yum -y install nc
RUN useradd testuser
COPY istio-tester /usr/bin
CMD /usr/bin/istio-tester