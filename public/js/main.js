const apresentacao = document.getElementById('apresentacao')
window.onload = () => {
	const dataAtual = new Date();
	const dataInicio = new Date('2021-05-01');
	const diferencaAnos = dataAtual.getFullYear() - dataInicio.getFullYear();
	apresentacao.innerText = apresentacao.innerText.replace('anos_inicio_programacao', diferencaAnos)
}
