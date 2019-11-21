FROM alpine
ADD shippydemo-srv /shippydemo-srv
ENTRYPOINT [ "/shippydemo-srv" ]
