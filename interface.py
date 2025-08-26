import subprocess
import tkinter as tk
from tkinter import messagebox

def calcular():
    num1 = entrada_num1.get()
    num2 = entrada_num2.get()
    op = entrada_op.get()

    try:
        resultado = subprocess.check_output(['./calc.exe', num1, num2, op], text=True)
        resultado = resultado.strip()
        label_resultado.config(text=f"Resultado: {resultado}")
    except Exception as e:
        messagebox.showerror("Erro", f"Erro ao chamar calculadora Go: {e}")

root = tk.Tk()
root.title("Calculadora Python chamando Go")

tk.Label(root, text="Número 1").grid(row=0, column=0)
entrada_num1 = tk.Entry(root)
entrada_num1.grid(row=0, column=1)

tk.Label(root, text="Número 2").grid(row=1, column=0)
entrada_num2 = tk.Entry(root)
entrada_num2.grid(row=1, column=1)

tk.Label(root, text="Operação (+, -, *, /)").grid(row=2, column=0)
entrada_op = tk.Entry(root)
entrada_op.grid(row=2, column=1)

btn = tk.Button(root, text="Calcular", command=calcular)
btn.grid(row=3, column=0, columnspan=2)

label_resultado = tk.Label(root, text="Resultado:")
label_resultado.grid(row=4, column=0, columnspan=2)

root.mainloop()
