 Implementations: 


 [Golang](go/README.md)



## How to test: 

run any server implementation and use this curl with a valid signature for integration env


```
curl --location --request PUT 'http://localhost:8080/webhook' \
--header 'Authorization: SHA256withRSA moZlAYs5TF/0ULF5H4hAQJbfGXQbgK7dEqgpPYbOnR0myHnLyeMGNSea3vNT9d5jiVQ326eWzzryKREvLqqIpvowhSjESfaQ0iBOAy31OZGrXooh/DQqfz3tI5Q/1JhQu3cpXjNVJyZ+2bZ63Ub+9dZucKjs5dcYmo6moB33uM+ErZtgix9q9brffNBIIjWMfH7+5wBuTdSysmbPjK6i9p8qJEoZtiJUVg3DGga/ykg5f/ncYCVBb+aNCDowLCb+UGwOtZHP7B8a+KkfHs6jcOvBqafy2TfZiRLVlILl0groyIxJAF3RHvk0F5F7H5Ziu6IHJHVIVsb/rnQ5MBfYpA==' \
--header 'Content-Type: application/json' \
--header 'Accept: application/json' \
--data-raw '{"paymentId":"4056d633-0d4c-4c11-aeb1-439022033d5a","merchantPaymentId":"3905","paymentStatus":"FAILED","errorDescription":null,"paymentRequest":{"amount":1900.25,"currency":"GBP","reference":"3905"},"paymentRefundData":null,"paymentMetadata":null}'

```
