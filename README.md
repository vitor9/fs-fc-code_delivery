# Full Stack-Full Cycle - Code Delivery / fs-fc-code_delivery
Repositorio dedicado a maratona Full Stack e Full Cycle

# O que sera desenvolvido neste projeto?
- Sistemas de entregas que permite visualizar em tempo real o veículo do entregador.
- Ha a possibilidade de multiplos entregadores simultaneos.
- Servico simulador que enviara a posicao em tempo real de cada entregador.
- Os dados de cada entrega, bem como as posicoes, serao armazenadas no ElasticSearch para futuras analises.

## Tecnologias a serem utilizadas:
## Simulador: GoLang.
        Possui um sistema que facilita trabalhar com multi-threading. Vou ter que enviar para o Kafka diversas posicoes de diversas corridas ao mesmo tempo, simultaneamente, e isso o multi-thread ajuda muito.

## BackEnd: Nest.js e Mongo.
        Nest.js roda em cima do node, e eh especifico em trabalhar com microsservicos. Vai ter bastante facilidade com websocket, e fazer comunicacao com o Kafka e o MongoDB
        MongoDB vai ser o nosso mecanismo nao-relacional, de persistencia de dados.

## FrontEnd: React.
        Vamos utilizar a Material.UI para deixar a interface mais agradavel e simples, para trabalharmos.

## Kafka & Kafka Connect.
        Kafka vai ser o nosso servidor de Streaming de dados que recebe as informacoes e consegue prover em tempo real.
        O Connect, vai se conectar no Kafka e ira jogar as informacoes em diversos sistemas.

## ElasticSearch e Kibana.
        O kibana eh um dashboard, que consegue acessar os dados do elasticsearch.

## Docker e Kubernetes
	Docker vai nos ajudar a montar o ambiente tranquilo, para subirmos todas estas tecnologias com facilidade, ao inves de instalar tecnologia por tecnologia
        Kubernetes vai nos auxiliar com a orquestracao dos containers do Docker, para que subirmos os containers com mais praticidade.

## Istio, Kiali, Prometheus e Grafana
        Istio, servico de service mesh, vai monitorar todas as comunicacoes dos nossos sistemas, para monitorarmos tudo que esta acontecendo entre os nossos servicos utilizando o Kiali.
        Grafana, Vai se conectar com o sistemas que vai coletar as metricas, utilizando o Prometheus.


