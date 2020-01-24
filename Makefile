GOPKG ?=	moul.io/lexique
DOCKER_IMAGE ?=	moul/lexique
GOBINS ?=	.
NPM_PACKAGES ?=	.

all: test install

-include rules.mk
