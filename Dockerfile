FROM python:2.7-slim
WORKDIR /app
ADD . /app
EXPOSE 5003
CMD ["/bin/sh", "-c", "python tracer.py > /dev/stdout 2>&1"]
