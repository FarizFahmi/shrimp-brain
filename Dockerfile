FROM golang:alpine AS builder

WORKDIR /usr/local/go/src/shrimp-brain

COPY ./go.mod .
COPY ./go.sum .
COPY ./main.go .

RUN go install github.com/air-verse/air@latest && go mod download

FROM alpine:3.14 AS final

WORKDIR /usr/local/go/src/shrimp-brain

# Install Go in the final stage
RUN apk add --no-cache curl bash
RUN curl -LO https://go.dev/dl/go1.23.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz && \
    rm go1.23.0.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

# # Install tzdata package
# RUN apk add --no-cache tzdata

# # Set the timezone (e.g., for UTC, America/New_York, etc.)
# ENV TZ=Asia/Jakarta

# # Copy timezone data
# RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && \
#     echo $TZ > /etc/timezone

# # Clean up tzdata package after setting the timezone
# RUN apk del tzdata

COPY --from=builder /go/bin/air /usr/local/bin/air
COPY --from=builder /usr/local/go/src/shrimp-brain /usr/local/go/src/shrimp-brain
COPY .air.toml /usr/local/go/src/shrimp-brain/.air.toml

CMD ["air", "-c", ".air.toml"]