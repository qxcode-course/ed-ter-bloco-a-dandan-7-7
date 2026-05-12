package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Editor struct {
	text   *List[*List[rune]] // a lista de linhas
	itLine *Node[*List[rune]] // iterador para a linha corrente
	itChar *Node[rune]        // iterador para o caracter do cursor
	screen tcell.Screen
	style  tcell.Style
}

func (e *Editor) InsertChar(r rune) {
	e.itChar = e.itLine.Value.Insert(e.itChar, r) // insere antes do elemento apontado pelo cursor
	e.itChar = e.itChar.Next()                    // move o cursor para próxima posição
}

func (e *Editor) KeyLeft() {
	if e.itChar != e.itLine.Value.Front() { // Se o cursor não está no início da linha
		e.itChar = e.itChar.Prev() // Move o cursor para a esquerda
		return
	}
	// Estamos no início da linha
	if e.itLine != e.text.Front() { // Se não está na primeira linha
		e.itLine = e.itLine.Prev()      // Atualiza iterador de linha para linha anterior
		e.itChar = e.itLine.Value.End() // Move o cursor para o final da linha
	}
}

func (e *Editor) KeyEnter() {
	
	nova := NewList[rune]()
	for e.itChar != e.itLine.Value.End() {
        nova.PushBack(e.itChar.Value)      // Salva o valor na nova lista
        e.itChar = e.itLine.Value.Erase(e.itChar) // Apaga da velha e itChar já vira o próximo nó
    }

	novaLinhaNode := e.text.Insert(e.itLine.Next(), nova)

	e.itLine = novaLinhaNode
	e.itChar = e.itLine.Value.Front()

	
}

func (e *Editor) KeyRight() {
	if e.itChar != e.itLine.Value.End() { 
		e.itChar = e.itChar.Next() 
		return
	}
	// Estamos no início da linha
	if e.itLine != e.text.Back() { 
		e.itLine = e.itLine.Next()      
		e.itChar = e.itLine.Value.Front() 
	}
}

func (e *Editor) KeyUp() {
	if e.itLine == e.text.Front(){
		return
	}
	
	ColunaAtual := e.itLine.Value.IndexOf(e.itChar)
	e.itLine = e.itLine.Prev()
	
	e.itChar = e.itLine.Value.Front()

	for i:= 0; i < ColunaAtual; i++{
		if e.itChar == e.itLine.Value.End(){
			break
			
		}
		e.itChar = e.itChar.Next()
	}
	
}

func (e *Editor) KeyDown() {

	if e.itLine == e.text.Back(){
		return
	}

	ColunaAtual := e.itLine.Value.IndexOf(e.itChar)
	e.itLine = e.itLine.Next()

	e.itChar = e.itLine.Value.Front()

	for i:= 0; i < ColunaAtual; i++{
		if e.itChar == e.itLine.Value.End(){
			break
			
		}
		e.itChar = e.itChar.Next()
	}
	

}

func (e *Editor) KeyBackspace() {
	
	if e.itChar != e.itLine.Value.Front() {
 
        e.itChar = e.itLine.Value.Erase(e.itChar.Prev())
        return 
    }

    if e.itLine == e.text.Front() {
        return
    }

    linhaAtual := e.itLine
    linhaDeCima := e.itLine.Prev()

    novoFocoCursor := linhaDeCima.Value.End()

    for eti := linhaAtual.Value.Front(); eti != linhaAtual.Value.End(); eti = eti.Next() {
        novoNo := linhaDeCima.Value.Insert(linhaDeCima.Value.End(), eti.Value)
        
        if eti == linhaAtual.Value.Front() {
            novoFocoCursor = novoNo
        }
    }

    e.itLine = linhaDeCima
    e.itChar = novoFocoCursor

    e.text.Erase(linhaAtual)


}

func (e *Editor) KeyDelete() {
	if e.itChar != e.itLine.Value.End() {
        
        e.itChar = e.itLine.Value.Erase(e.itChar)
        return 
    }

	if e.itLine == e.text.Back() {
        return
    }

	linhaDeBaixo := e.itLine.Next()
    novoFocoCursor := e.itLine.Value.End()

    for eti := linhaDeBaixo.Value.Front(); eti != linhaDeBaixo.Value.End(); eti = eti.Next() {
        novoNo := e.itLine.Value.Insert(e.itLine.Value.End(), eti.Value)
        if eti == linhaDeBaixo.Value.Front() {
            novoFocoCursor = novoNo
        }
    }

    e.itChar = novoFocoCursor
    e.text.Erase(linhaDeBaixo)
}	
		


func main() { // Texto inicial e posição do cursor
	editor := NewEditor()
	editor.Draw()
	editor.MainLoop()
	defer editor.screen.Fini() // Encerra a tela ao sair
}

func (e *Editor) MainLoop() {
	for {
		ev := e.screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc, tcell.KeyCtrlC:
				return
			case tcell.KeyEnter:
				e.KeyEnter()
			case tcell.KeyLeft:
				e.KeyLeft()
			case tcell.KeyRight:
				e.KeyRight()
			case tcell.KeyUp:
				e.KeyUp()
			case tcell.KeyDown:
				e.KeyDown()
			case tcell.KeyBackspace, tcell.KeyBackspace2:
				e.KeyBackspace()
			case tcell.KeyDelete:
				e.KeyDelete()
			default:
				if ev.Rune() != 0 {
					e.InsertChar(ev.Rune())
				}
			}
			e.Draw()
		case *tcell.EventResize:
			e.screen.Sync()
			e.Draw()
		}
	}
}

func NewEditor() *Editor {
	e := &Editor{}
	// Inicializa a tela
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Printf("erro ao criar a tela: %v", err)
	}
	if err := screen.Init(); err != nil {
		fmt.Printf("erro ao iniciar a tela: %v", err)
	}
	e.screen = screen
	e.text = NewList[*List[rune]]()
	e.text.PushBack(NewList[rune]())
	e.itLine = e.text.Front()
	e.itChar = e.itLine.Value.Back()
	// Define o estilo do texto (branco com fundo preto)
	e.style = tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorBlack)

	// Limpa a tela e define o estilo base
	e.screen.SetStyle(e.style)
	e.screen.Clear()
	return e
}

func (e *Editor) Draw() {
	e.screen.Clear()
	x := 0
	y := 0
	for line := e.text.Front(); line != e.text.End(); line = line.Next() {
		for char := line.Value.Front(); ; char = char.Next() {
			data := char.Value
			if char == line.Value.End() {
				data = '↲'
			}
			if data == ' ' {
				data = '·'
			}
			if char == e.itChar {
				e.screen.SetContent(x, y, data, nil, e.style.Reverse(true))
			} else {
				e.screen.SetContent(x, y, data, nil, e.style)
			}
			x++
			if char == line.Value.End() {
				break
			}
		}
		y++
		x = 0
	}
	e.screen.Show()
}
