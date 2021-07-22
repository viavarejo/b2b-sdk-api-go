package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/viavarejo/b2b-sdk-api-go/src/api"
	"github.com/viavarejo/b2b-sdk-api-go/src/model/request"
	"github.com/viavarejo/b2b-sdk-api-go/src/model/response"
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

var idSkuCriacaoPedido int64 = 8935731
var idSkuCriacaoPedidoComCartao int64 = 9342200
var IdLojista int64 = 15
var IdCampanha int64 = 5940
var cnpj string = "57.822.975/0001-12"
var cep string = "01525000"
var cpfDestinatario string = "435.375.660-50"
var pedidoHelper DadosPedidoHelper
var pedidoComCartaoHelper DadosPedidoHelper

func TestPostCalcularCarrinhoParaCriacaoPedido_1(t *testing.T) {
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

}

func TestPostCalcularCarrinhoParaCriacaoPedidoComCartao_2(t *testing.T) {
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
}

func TestPostCriarPedido_3(t *testing.T) {
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
		t.Error("Test failed-1")
	}

	pedidoHelper.IdPedido = dto.Data.CodigoPedido
	pedidoHelper.IdPedidoParceiro = dto.Data.PedidoParceiro

}

func TestPostCriarPedidoPagCartao_4(t *testing.T) {
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

	//Adicionar criptografia
	dadosCartaoCredito.Nome = "nome"
	dadosCartaoCredito.Numero = "nome"
	dadosCartaoCredito.CodigoVerificador = "nome"
	dadosCartaoCredito.ValidadeAno = "nome"
	dadosCartaoCredito.ValidadeMes = "nome"
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

}

func TestPatchPedidosCancelamento_5(t *testing.T) {
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

}

func TestPatchPedidosConfirmacao_6(t *testing.T) {
	confirmacao := request.ConfirmacaoReq{}

	confirmacao.IDCampanha = IdCampanha
	confirmacao.IDPedidoParceiro = pedidoComCartaoHelper.IdPedidoParceiro
	confirmacao.Confirmado = true

	dto := api.PatchPedidosCancelamentoConfirmacao(fmt.Sprint(pedidoComCartaoHelper.IdPedido), confirmacao)

	if !dto.Data.PedidoConfirmado {
		t.Error("Test failed-1")
	}

}

func TestGetDadosPedidoParceiro_7(t *testing.T) {

	dto := api.GetDadosPedidoParceiro(fmt.Sprint(pedidoHelper.IdPedido), cnpj, fmt.Sprint(IdCampanha), fmt.Sprint(pedidoHelper.IdPedidoParceiro), "")

	if !(pedidoHelper.IdPedido == dto.Data.Pedido.CodigoPedido) {
		t.Error("Test failed-1")
	}
}

func TestGetNotaFiscalPedidoPdf_8(t *testing.T) {
	dto := api.GetNotaFiscalPedido("247473612", "91712686", "PDF")

	if &dto == nil {
		t.Error("Test failed-1")
	}
}

func TestGetDadosPedidoParceiroFail_9(t *testing.T) {
	dto := api.GetDadosPedidoParceiro(fmt.Sprint(pedidoHelper.IdPedido), cnpj, fmt.Sprint(IdCampanha), fmt.Sprint(pedidoHelper.IdPedidoParceiro), "")

	if !("400" == dto.Error.Code) {
		t.Error("Test failed-1")
	}

}

func TestPostCalcularCarrinhoParaCriacaoPedidoFail_10(t *testing.T) {
	request_dto := request.PedidoCarrinho{}
	request_dto.Cnpj = cnpj

	dto := api.PostCalcularCarrinho(request_dto)

	if &dto != nil {
		t.Error("Test failed-1")
	}
}

func TestPatchPedidosFail_11(t *testing.T) {
	confirmacao := request.ConfirmacaoReq{}

	confirmacao.IDCampanha = IdCampanha
	confirmacao.IDPedidoParceiro = pedidoHelper.IdPedidoParceiro
	confirmacao.Cancelado = true

	dto := api.PatchPedidosCancelamentoConfirmacao("", confirmacao)

	if !dto.Data.PedidoConfirmado {
		t.Error("Test failed-1")
	}
}

func TestPatchPedidosConfirmacaoFail_12(t *testing.T) {
	confirmacao := request.ConfirmacaoReq{}

	dto := api.PatchPedidosCancelamentoConfirmacao(fmt.Sprint(pedidoComCartaoHelper.IdPedido), confirmacao)

	if !dto.Data.PedidoConfirmado {
		t.Error("Test failed-1")
	}

}

func TestGetNotaFiscalPedidoFail_13(t *testing.T) {
	dto := api.GetNotaFiscalPedido("247473612", "91712686", "PDF")

	if &dto == nil {
		t.Error("Test failed-1")
	}
}

func TestPostCriarPedidoFail_14(t *testing.T) {
	criacao := request.CriacaoPedidoReq{}
	dto := api.PostCriarPedido(criacao)
	if &dto == nil {
		t.Error("Test failed-1")
	}
}

func preparaPedido(calculo response.CalculoCarrinho) DadosPedidoHelper {
	helper := DadosPedidoHelper{}
	helper.IdSku = calculo.Data.Produtos[0].IDSku
	helper.precoVenda = calculo.Data.Produtos[0].ValorUnitario
	helper.ValorFrete = calculo.Data.Produtos[0].ValorTotalFrete
	return helper
}

func geraIdPedidoParceiro() int64 {
	return rand.Int63n(979899)
}
