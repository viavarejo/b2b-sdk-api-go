package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/viavarejo/b2b-sdk-api-go/src/api"
	"github.com/viavarejo/b2b-sdk-api-go/src/model/request"
	"github.com/viavarejo/b2b-sdk-api-go/src/model/response"
	"github.com/viavarejo/b2b-sdk-api-go/src/security"
)

type DadosPedidoHelper struct {
	IdPedido         int64
	IdPedidoParceiro int64
	IdSku            int64
	ValorFrete       float64
	precoVenda       float64
}
type DadosCartaoHelper struct {
	Nome              string
	Numero            string
	CodigoVerificador string
	AnoValidade       string
	MesValidade       string
}

/**
 * Os testes para as URI's dos Pedidos do B2B.
 * É importante que os metodos sejam executados na ordem estabelecida, pois
 * alguns metodos de testes possuem dependencia dos resultados dos anteriores.
 *
 */

var idSkuCriacaoPedido int64 = 8935731
var idSkuCriacaoPedidoComCartao int64 = 9342200
var IdLojista int64 = 15
var IdCampanha int64 = 5940
var cnpj string = "57.822.975/0001-12"
var cep string = "01525000"
var cpfDestinatario string = "435.375.660-50"
var pedidoHelper DadosPedidoHelper
var pedidoComCartaoHelper DadosPedidoHelper

func Test_1PostCalcularCarrinhoParaCriacaoPedido(t *testing.T) {
	t.Run("Deveria calcular o carrinho para criação do pedido", func(t *testing.T) {
		produto := request.ProdutoPedidoCarrinho{}
		produto.Codigo = idSkuCriacaoPedido
		produto.Quantidade = 1
		produto.IDLojista = IdLojista

		pedido := request.PedidoCarrinho{}
		pedido.IDCampanha = IdCampanha
		pedido.Cnpj = cnpj
		pedido.Cep = cep
		pedido.Produtos = []request.ProdutoPedidoCarrinho{produto}
		dto := api.PostCalcularCarrinho(pedido)

		if &dto.Data == nil {
			t.Error("Test failed-1")
		}
		if dto.Data.ValorFrete <= 0.0 {
			t.Error("Test failed-2")
		}
		if dto.Data.ValorTotaldoPedido <= 0.0 {
			t.Error("Test failed-3")
		}
		if dto.Data.Produtos[0].ValorTotalFrete <= 0.0 {
			t.Error("Test failed-4")
		}
		pedidoHelper = preparaPedido(dto)
	})
}

func Test_2PostCalcularCarrinhoParaCriacaoPedidoComCartao(t *testing.T) {
	t.Run("Deveria calcular o carrinho para criação do pedido com cartão", func(t *testing.T) {
		var produto request.ProdutoPedidoCarrinho
		produto.Codigo = idSkuCriacaoPedidoComCartao
		produto.Quantidade = 1
		produto.IDLojista = IdLojista

		var pedido request.PedidoCarrinho
		pedido.IDCampanha = IdCampanha
		pedido.Cnpj = cnpj
		pedido.Cep = cep
		pedido.Produtos = []request.ProdutoPedidoCarrinho{produto}
		dto := api.PostCalcularCarrinho(pedido)

		if &dto.Data == nil {
			t.Error("Test failed-1")
		}
		if dto.Data.ValorFrete <= 0.0 {
			t.Error("Test failed-2")
		}
		if dto.Data.ValorTotaldoPedido <= 0.0 {
			t.Error("Test failed-3")
		}
		if dto.Data.Produtos[0].ValorTotalFrete <= 0.0 {
			t.Error("Test failed-4")
		}
		pedidoComCartaoHelper = preparaPedido(dto)
	})
}

func Test_3PostCriarPedido(t *testing.T) {
	t.Run("Deveria criar o pedido", func(t *testing.T) {
		request_dto := request.CriacaoPedidoReq{}
		produto := request.ProdutoCriacaoPedido{}
		enderecoEntrega := request.Endereco{}
		destinatario := request.Destinatario{}
		dadosEntrega := request.DadosEntrega{}

		produto.IDLojista = IdLojista
		produto.Codigo = pedidoHelper.IdSku
		produto.Quantidade = 1
		produto.PrecoVenda = pedidoHelper.precoVenda
		produto.Premio = 0

		enderecoEntrega.Cep = "01525-000"
		enderecoEntrega.Estado = "SP"
		enderecoEntrega.Logradouro = "rua da se"
		enderecoEntrega.Cidade = "São Paulo"
		enderecoEntrega.Numero = 63
		enderecoEntrega.Referencia = "teste"
		enderecoEntrega.Bairro = "bairro se"
		enderecoEntrega.Complemento = "teste"
		enderecoEntrega.Telefone = "22333333"
		enderecoEntrega.Telefone2 = "22333335"
		enderecoEntrega.Telefone3 = "22333336"

		destinatario.Nome = "teste"
		destinatario.CpfCnpj = cpfDestinatario
		destinatario.Email = "teste@teste.com"

		dadosEntrega.ValorFrete = pedidoHelper.ValorFrete

		request_dto.Produtos = []request.ProdutoCriacaoPedido{produto}
		request_dto.EnderecoEntrega = enderecoEntrega
		request_dto.Destinatario = destinatario
		request_dto.DadosEntrega = dadosEntrega
		request_dto.Campanha = IdCampanha
		request_dto.Cnpj = cnpj
		request_dto.PedidoParceiro = geraIdPedidoParceiro()
		request_dto.ValorFrete = pedidoHelper.ValorFrete
		request_dto.AguardarConfirmacao = true
		request_dto.OptantePeloSimples = true

		dto := api.PostCriarPedido(request_dto)
		valorTotal := pedidoHelper.ValorFrete + pedidoHelper.precoVenda

		if valorTotal != dto.Data.ValorTotalPedido {
			fmt.Println("valorTotal: " + fmt.Sprintf("%f", valorTotal))
			fmt.Println("dto.Data.ValorTotalPedido: " + fmt.Sprintf("%f", dto.Data.ValorTotalPedido))
			t.Error("Test failed-1")
		}

		pedidoHelper.IdPedido = dto.Data.CodigoPedido
		pedidoHelper.IdPedidoParceiro = dto.Data.PedidoParceiro
	})

}

func Test_4PostCriarPedidoPagCartao(t *testing.T) {
	t.Run("Deveria calcular o carrinho para criação do pedido com cartão", func(t *testing.T) {
		request_dto := request.CriacaoPedidoReq{}
		produto := request.ProdutoCriacaoPedido{}
		enderecoEntrega := request.Endereco{}
		destinatario := request.Destinatario{}
		dadosEntrega := request.DadosEntrega{}
		pagtosComplementares := request.PagtosComplementare{}
		dadosCartaoCredito := request.DadosCartaoCredito{}
		dadosCartaoCreditoValidacao := request.DadosCartaoCreditoValidacao{}

		produto.IDLojista = IdLojista
		produto.Codigo = pedidoComCartaoHelper.IdSku
		produto.Quantidade = 1
		produto.PrecoVenda = pedidoComCartaoHelper.precoVenda
		produto.Premio = 0

		enderecoEntrega.Cep = "01525-000"
		enderecoEntrega.Estado = "SP"
		enderecoEntrega.Logradouro = "rua da se"
		enderecoEntrega.Cidade = "São Paulo"
		enderecoEntrega.Numero = 63
		enderecoEntrega.Referencia = "teste"
		enderecoEntrega.Bairro = "bairro se"
		enderecoEntrega.Complemento = "teste"
		enderecoEntrega.Telefone = "22333333"
		enderecoEntrega.Telefone2 = "22333335"
		enderecoEntrega.Telefone3 = "22333336"

		destinatario.Nome = "teste"
		destinatario.CpfCnpj = cpfDestinatario
		destinatario.Email = "teste@teste.com"

		dadosEntrega.ValorFrete = pedidoComCartaoHelper.ValorFrete

		var chaveDto response.ChaveDTO = api.GetChave()

		//Adicionar criptografia
		dadosCartaoCredito.Nome = security.Encrypt(chaveDto.Data.ChavePublica, "Jose da Silva")
		dadosCartaoCredito.Numero = security.Encrypt(chaveDto.Data.ChavePublica, "5155901222280001")
		dadosCartaoCredito.CodigoVerificador = security.Encrypt(chaveDto.Data.ChavePublica, "1234")
		dadosCartaoCredito.ValidadeAno = security.Encrypt(chaveDto.Data.ChavePublica, "2045")
		dadosCartaoCredito.ValidadeMes = security.Encrypt(chaveDto.Data.ChavePublica, "12")
		dadosCartaoCredito.QuantidadeParcelas = 1

		dadosCartaoCreditoValidacao.Nome = "Jose da Silva"
		dadosCartaoCreditoValidacao.NumeroMascarado = "515590XXXXXX0001"
		dadosCartaoCreditoValidacao.QtCaracteresCodigoVerificador = "4"
		dadosCartaoCreditoValidacao.ValidadeAno = "2045"
		dadosCartaoCreditoValidacao.ValidadeMes = "12"

		pagtosComplementares.IDFormaPagamento = 3
		pagtosComplementares.DadosCartaoCredito = dadosCartaoCredito
		pagtosComplementares.DadosCartaoCreditoValidacao = dadosCartaoCreditoValidacao
		pagtosComplementares.ValorComplementar = 30.0
		pagtosComplementares.ValorComplementarCOMJuros = 30.0

		request_dto.Campanha = IdCampanha
		request_dto.Cnpj = cnpj
		request_dto.PedidoParceiro = geraIdPedidoParceiro()
		request_dto.ValorFrete = pedidoComCartaoHelper.ValorFrete
		request_dto.OptantePeloSimples = true
		request_dto.AguardarConfirmacao = true
		request_dto.PossuiPagtoComplementar = true
		request_dto.Produtos = []request.ProdutoCriacaoPedido{produto}
		request_dto.EnderecoEntrega = enderecoEntrega
		request_dto.Destinatario = destinatario
		request_dto.DadosEntrega = dadosEntrega
		request_dto.EnderecoCobranca = enderecoEntrega
		request_dto.PagtosComplementares = []request.PagtosComplementare{pagtosComplementares}
		request_dto.ValorTotalPedido = pedidoComCartaoHelper.ValorFrete + pedidoComCartaoHelper.precoVenda
		request_dto.ValorTotalComplementar = 30.0
		request_dto.ValorTotalComplementarCOMJuros = 30.0

		dto := api.PostCriarPedido(request_dto)
		valorTotal := pedidoComCartaoHelper.ValorFrete + pedidoComCartaoHelper.precoVenda

		if valorTotal != dto.Data.ValorTotalPedido {
			t.Error("Test failed-1")
		}

		pedidoComCartaoHelper.IdPedido = dto.Data.CodigoPedido
		pedidoComCartaoHelper.IdPedidoParceiro = dto.Data.PedidoParceiro
	})

}

func Test_5PatchPedidosCancelamento(t *testing.T) {
	t.Run("Deveria fazer o cancelamento do pedido", func(t *testing.T) {
		confirmacao := request.ConfirmacaoReq{}

		confirmacao.IDCampanha = IdCampanha
		confirmacao.IDPedidoParceiro = pedidoHelper.IdPedidoParceiro
		confirmacao.Cancelado = true
		confirmacao.Confirmado = false
		confirmacao.IDPedidoMktplc = "1-01"
		confirmacao.MotivoCancelamento = "teste"
		confirmacao.Parceiro = "BANCO INTER"
		val := strconv.Itoa(int(pedidoHelper.IdPedido))
		dto := api.PatchPedidosCancelamentoConfirmacao(val, confirmacao)

		if !dto.Data.PedidoCancelado {
			t.Error("Test failed-1")
		}
	})

}

func Test_6PatchPedidosConfirmacao(t *testing.T) {
	t.Run("Deveria confirmar o pedido", func(t *testing.T) {
		confirmacao := request.ConfirmacaoReq{}

		confirmacao.IDCampanha = IdCampanha
		confirmacao.IDPedidoParceiro = pedidoComCartaoHelper.IdPedidoParceiro
		confirmacao.Confirmado = true

		dto := api.PatchPedidosCancelamentoConfirmacao(fmt.Sprint(pedidoComCartaoHelper.IdPedido), confirmacao)

		if !dto.Data.PedidoConfirmado {
			t.Error("Test failed-1")
		}
	})

}

func Test_07GetDadosPedidoParceiro(t *testing.T) {
	t.Run("Deveria buscar os dados do pedido parceiro", func(t *testing.T) {
		dto := api.GetDadosPedidoParceiro(fmt.Sprint(pedidoHelper.IdPedido), cnpj, fmt.Sprint(IdCampanha), fmt.Sprint(pedidoHelper.IdPedidoParceiro), "")

		if !(pedidoHelper.IdPedido == dto.Data.Pedido.CodigoPedido) {
			t.Error("Test failed-1")
		}
	})
}

func Test_8GetNotaFiscalPedidoPdf(t *testing.T) {
	t.Run("Deveria retornar a nota fiscal no formato de pdf", func(t *testing.T) {
		_, resp := api.GetNotaFiscalPedido("247473612", "91712686", "XML")

		if &resp.Body == nil {
			t.Error("Test failed-1")
		}
	})
}

func Test_9GetNotaFiscalPedidoXml(t *testing.T) {
	t.Run("Deveria retornar a nota fiscal no formato de xml", func(t *testing.T) {
		_, resp := api.GetNotaFiscalPedido("247473612", "91712686", "PDF")

		if &resp.Body == nil {
			t.Error("Test failed-1")
		}
	})
}

func Test_10GetDadosPedidoParceiroFail(t *testing.T) {
	t.Run("Falhar na busca de dados do pedido parceiro", func(t *testing.T) {
		dto := api.GetDadosPedidoParceiro(fmt.Sprint(pedidoHelper.IdPedido), cnpj, "", "", "")

		if !("400" == dto.Error.Code) {
			t.Error("Test failed-1")
		}
	})
}

func Test_11PostCalcularCarrinhoParaCriacaoPedidoFail(t *testing.T) {
	t.Run("Falhar ao calcular o carrinho", func(t *testing.T) {
		request_dto := request.PedidoCarrinho{}
		request_dto.Cnpj = cnpj

		dto := api.PostCalcularCarrinho(request_dto)

		if !(&dto == nil) {
			t.Error("Test failed-1")
		}
	})
}

func Test_12PatchPedidosFail(t *testing.T) {
	t.Run("Falhar na confirmação do pedido", func(t *testing.T) {
		confirmacao := request.ConfirmacaoReq{}

		confirmacao.IDCampanha = IdCampanha

		dto := api.PatchPedidosCancelamentoConfirmacao("123", confirmacao)

		if dto.Error.Code == "400" {
			t.Error("Test failed-1")
		}
	})
}

func Test_13PatchPedidosConfirmacaoFail(t *testing.T) {
	t.Run("Falhar na confirmação do pedido 2", func(t *testing.T) {
		confirmacao := request.ConfirmacaoReq{}
		dto := api.PatchPedidosCancelamentoConfirmacao(fmt.Sprint(pedidoComCartaoHelper.IdPedido), confirmacao)
		if dto.Error.Code == "400" {
			t.Error("Test failed-1")
		}
	})
}

func Test_14GetNotaFiscalPedidoFail(t *testing.T) {
	t.Run("Falhar na busca pela nota fiscal", func(t *testing.T) {
		_, resp := api.GetNotaFiscalPedido("247473612", "91712686", "PDF")
		if &resp == nil {
			t.Error("Test failed-1")
		}
	})
}

func Test_15PostCriarPedidoFail(t *testing.T) {
	t.Run("Falhar na criacao do pedido", func(t *testing.T) {
		criacao := request.CriacaoPedidoReq{}
		dto := api.PostCriarPedido(criacao)
		if &dto == nil {
			t.Error("Test failed-1")
		}
	})
}

func preparaPedido(calculo response.CalculoCarrinho) DadosPedidoHelper {
	helper := DadosPedidoHelper{}
	helper.IdSku = calculo.Data.Produtos[0].IDSku
	helper.precoVenda = calculo.Data.Produtos[0].ValorUnitario
	helper.ValorFrete = calculo.Data.Produtos[0].ValorTotalFrete
	return helper
}

func geraIdPedidoParceiro() int64 {
	return rand.Int63n(85812)
}
