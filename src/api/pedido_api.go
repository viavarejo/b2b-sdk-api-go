package api

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	req "github.com/viavarejo/b2b-sdk-api-go/src/model/request"
	resp "github.com/viavarejo/b2b-sdk-api-go/src/model/response"
	"github.com/viavarejo/b2b-sdk-api-go/src/service"
)

func PostCalcularCarrinho(pedido req.PedidoCarrinho) resp.CalculoCarrinho {
	body, _ := pedido.Marshal()
	response := service.Post(body, "/pedidos/carrinho")
	dto, _ := resp.UnmarshalCalculoCarrinho([]byte(response))
	return dto
}

func GetDadosPedidoParceiro(idCompra string, cnpj string, idCampanha string, idPedidoParceiro string, idPedidoMktplc string) resp.PedidoParceiroDados {
	var queryParams = make(map[string]interface{})
	if len(cnpj) > 0 {
		queryParams["request.cnpj"] = cnpj
	}
	if len(cnpj) > 0 {
		queryParams["request.idCampanha"] = idCampanha
	}
	if len(cnpj) > 0 {
		queryParams["request.idPedidoParceiro"] = idPedidoParceiro
	}
	if len(cnpj) > 0 {
		queryParams["request.idPedidoMktplc"] = idPedidoMktplc
	}

	response := service.Get("/pedidos/"+idCompra, queryParams)
	pedido, _ := resp.UnmarshalPedidoParceiroDados([]byte(response))

	return pedido
}

func PatchPedidosCancelamentoConfirmacao(idCompra string, confirmacao req.ConfirmacaoReq) resp.ConfirmacaoResp {
	body, _ := confirmacao.Marshal()
	response := service.Patch("/pedidos/"+idCompra, body)
	dto, _ := resp.UnmarshalConfirmacaoResp([]byte(response))
	return dto
}

func GetNotaFiscalPedido(idCompra string, idCompraEntrega string, formato string) ([]byte, *http.Response) {
	urlPath := "/pedidos/" + idCompra + "/entregas/" + idCompraEntrega + "/nfe/" + formato

	file, resp := service.DownloadFile(urlPath)
	// Create the file
	out, err := os.Create(idCompra + "_" + idCompraEntrega + "." + strings.ToLower(formato))
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// Write the body to file
	//_, err = io.Copy(out, resp.Body)
	wfile, err := out.Write(file)
	if err != nil {
		panic(err)
	}
	fmt.Printf("wrote %d bytes\n", wfile)

	return file, resp
}

func PostCriarPedido(pedido req.CriacaoPedidoReq) resp.CriacaoPedidoResp {
	body, _ := pedido.Marshal()
	response := service.Post(body, "/pedidos")
	dto, _ := resp.UnmarshalCriacaoPedidoResp([]byte(response))
	return dto
}
