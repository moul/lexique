GOPKG ?=	moul.io/lexique
DOCKER_IMAGE ?=	moul/lexique
GOBINS ?=	.
NPM_PACKAGES ?=	.
LEXIQUE_ID ?= 383

PRE_INSTALL_STEPS += lexique.tsv

-include rules.mk

lexique.tsv:
	wget http://www.lexique.org/databases/Lexique$(LEXIQUE_ID)/Lexique$(LEXIQUE_ID).zip
	unzip Lexique$(LEXIQUE_ID).zip Lexique$(LEXIQUE_ID).tsv
	rm -f Lexique$(LEXIQUE_ID).zip
	mv Lexique$(LEXIQUE_ID).tsv lexique.tsv
