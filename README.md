<h3>Client-Server-Ethos-Go</h3>

<p>To develop two programs namely bankingserver.go, bankingclient.go and compile them for Ethos, and run them on Ethos operating system.</p>

<h3>Files in the project:</h3>
<ol>
<li><b>bankingclient.go</b> - the client file which consumes RPC's defined in server.</li>
<li><b>bankingserver.go</b> - the server file contains the definations of the 4 different RPC's namely balance, deposit, withdraw, transfer.</li>
<li><b>myRpc.t</b> - this is a interface file for RPC methods </li>
<li><b>MakeFile</b> - this is the build file</li>
<li><b>module.mk</b> - file with all definations of client, server and interface</li>
</ol>

<h3>Steps to follow to run the project:</h3>
<ol>
<li>Copy the submission files into ethos environment into a folder - let's say "homework"</li>

```
cd homework
```

<li>Let's kill if any process is running, by using the below command</li>

```
sudo -E ethosKillAll
```

<li>Now let's compile using the below command</li>

```
make && make clean && make install
```

<li>Now let's run the server</li>

```
cd client && sudo -E ethosRun
```

<li>Open another terminal and let's run the client</li>

```
etAl client.ethos
```
<li>If you want to launch multiple clients, after the above step go into client directory of where the project is compiled with another user's credentials and run the following command.</li>

```
et client.ethos
```

<li>You will be prompted with /user/me now enter the client file name. If multiple users are launched enter the client file name.</li>

```
bankingclient
```

<li>Now we can check the logs in the client directory and enter the below command</li>

```
ethosLog .
```

<li>The logs are attached in the submission folder for reference.</li>

<li>Below are few logs:</li>

```
1680031576.121956013 [bankingclient.O] ********************* Entering Client ************************
1680031576.129838092 [bankingclient.O] The current user is me
1680031576.137648758 [bankingclient.O] ********************* Entering Deposit ************************
1680031576.174405261 [bankingclient.O] Received deposit reply: The deposit of 60 is done and the balance is 1060
1680031576.186035490 [bankingclient.O] ********************* Exiting Deposit ************************
1680031576.193879419 [bankingclient.O] ********************* Entering Withdraw ************************
1680031576.201719809 [bankingclient.O] Received withdraw reply: The withdrwal of 60 is done and the balance is 1000
1680031576.217372095 [bankingclient.O] ********************* Exiting Withdraw ************************
1680031576.219488411 [bankingclient.O] ********************* Entering Transfer ************************
1680031576.234135421 [bankingclient.O] The transfer status: The amount of 10 has been transfered to bennett from me
1680031576.234841752 [bankingclient.O] ********************* Exiting Transfer ************************
1680031576.249986170 [bankingclient.O] ********************* Entering Balance ************************
1680031576.254712687 [bankingclient.O] The Balance in the account is: 990
1680031576.255573684 [bankingclient.O] ********************* Exiting Balance ************************
1680031576.256583367 [bankingclient.O] ********************* Exiting Client ************************
1680031584.096997616 [bankingclient.O] ********************* Entering Client ************************
1680031584.102720323 [bankingclient.O] The current user is jlong
1680031584.103245500 [bankingclient.O] ********************* Entering Deposit ************************
1680031584.105897075 [bankingclient.O] Received deposit reply: The deposit of 60 is done and the balance is 560
1680031584.107056155 [bankingclient.O] ********************* Exiting Deposit ************************
1680031584.108491831 [bankingclient.O] ********************* Entering Withdraw ************************
1680031584.112432916 [bankingclient.O] Received withdraw reply: The withdrwal of 60 is done and the balance is 500
1680031584.113271900 [bankingclient.O] ********************* Exiting Withdraw ************************
1680031584.116531134 [bankingclient.O] ********************* Entering Transfer ************************
1680031584.118371076 [bankingclient.O] The transfer status: The amount of 10 has been transfered to bennett from jlong
1680031584.119627096 [bankingclient.O] ********************* Exiting Transfer ************************
1680031584.120851641 [bankingclient.O] ********************* Entering Balance ************************
1680031584.125172627 [bankingclient.O] The Balance in the account is: 490
1680031584.126523882 [bankingclient.O] ********************* Exiting Balance ************************
1680031584.127529179 [bankingclient.O] ********************* Exiting Client ************************
1680031589.775942603 [bankingclient.O] ********************* Entering Client ************************
1680031589.787710136 [bankingclient.O] The current user is bennett
1680031589.788284746 [bankingclient.O] ********************* Entering Deposit ************************
1680031589.799209484 [bankingclient.O] Received deposit reply: The deposit of 60 is done and the balance is 380
1680031589.808034887 [bankingclient.O] ********************* Exiting Deposit ************************
1680031589.815420066 [bankingclient.O] ********************* Entering Withdraw ************************
1680031589.832679522 [bankingclient.O] Received withdraw reply: The withdrwal of 60 is done and the balance is 320
1680031589.846833350 [bankingclient.O] ********************* Exiting Withdraw ************************
1680031589.863786702 [bankingclient.O] ********************* Entering Transfer ************************
1680031589.865543216 [bankingclient.O] The transfer status: The amount of 10 has been transfered to bennett from bennett
1680031589.879171930 [bankingclient.O] ********************* Exiting Transfer ************************
1680031589.886736915 [bankingclient.O] ********************* Entering Balance ************************
1680031589.911536199 [bankingclient.O] The Balance in the account is: 320
1680031589.927089012 [bankingclient.O] ********************* Exiting Balance ************************
1680031589.938792720 [bankingclient.O] ********************* Exiting Client ************************
```

</ol>

<h3>References:</h3>
<ol>
<li>Syncclient example code provided on box</li>
<li>goOnEthos pdf manual</li>
</ol>
