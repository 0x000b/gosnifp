<h1 align="center">
<img src="./assets/icon.png" width=200>
    
Gosnifp


</h1>

<h4 align="center">

[Português](./README.md)

[English](./ENGLISH.md)

</h4>

A network sniffer focused on DNS communications.


<img src="./assets/initial.png" width=860>

## Installation

If you want all the project files, make sure you have `Go` installed and simply `clone` the project.

```bash
git clone https://github.com/0x000b/gosnifp.git
cd gosnifp/
```
In the folder, `run`:
```bash
go build cmd/gosnifp/main.go
```
Once you have obtained the binary file, just use `sudo` to allow the sniffer to read your interface:
```bash
sudo ./gosnifp [INTERFACE]
```

## Diagram

The diagram tries to represent the basic functioning of the sniffer. Incoming DNS packets pass through two main processes: Sniffer and Preprocessor. The `Sniffer` is actually the part of the software that actively captures packets on the network and sends them for processing. The `Preprocessor` is responsible for capturing the important data from the packet, creating an object relating to the packet and sending it to the monitors (`Monitor Logging`) and loggers who will create the `Log` files.

<img src="./assets/arch.png">

## FAQ 

<details>
<summary>
     What is Passive DNS?
</summary>

<p>
Passive DNS is the act of collecting and storing DNS communication for future analysis, i.e. it consists of saving DNS request and response data.
</p>
</details>


<details>
<summary>
    What is a Sniffer?
</summary>

<p>
A sniffer is a program or piece of hardware that intercepts, visualizes and records network traffic data.
</p>
</details>

<details>
<summary>
Can I contribute to the project?
</summary>

<p>
Any contribution is accepted and welcome, the project is initially public and serves as a basis for network studies, but can be used for real cases.
</p>
</details>