## 1. Quadratic nonresidue
### Completeness
该问题等价于证明 $QR(m,s^2x)=0$ 且 $QR(m,s^2)=1$ 。  
(1) 因为 $s^2\equiv s^2(mod m)$ ，所以 $QR(m,s^2)=1$ 是显然的;  
(2) 假设 $QR(m,s^2x)=1$ ，那么必然存在一个 $r \in \mathbb{Z}_m$ 使得 $r^2\equiv s^2x(mod m)$ ，又因为 s,m 互素，则必存在 $s^{-1}$ 使得 $s\*s^{-1}\equiv 1(mod m)$ ，结合假设可得 $(r*s^{-1})^2\equiv x(mod m)$，因此存在 $t=r\*s^{-1}$ 使得 $t^2\equiv x(mod m)$ ，与 $QR(m,x)=0$ 矛盾。因此 $QR(m,s^2x)=0$ 得证。
### Soundness
若 $QR(m,x)=1$ ，则存在 $r \in \mathbb{Z}_m$ 使得 $r^2\equiv x(mod m)$ ，进一步有 $(r\*s)^2\equiv s^2x(mod m)$，因此 $QR(m,s^2x)=1\ 且\ QR(m,s^2)=1$，因此 prover 无法区分 $s^2x$ 和 $s^2$，所以 prover 只有 50% 的概率答对，问题得证。

## 2. Quadratic residue
### Completeness:
该问题等价于证明 $xt^2\equiv xt^2(mod m)$ 且 $s^2t^2\equiv xt^2(mod m)$ 。  
对于前者这是显然的；  
对于后者，因为 $s^2\equiv x(mod m)$ ，因此 $s^2t^2\equiv xt^2(mod m)$ 。
### Soundness
易知 b=0 时 prover 总能答对，因此原问题等价于 b=1 时 prover 不可能答对，反证法易证。  
假设 b=1 时 prover 可以给出一个 u 使得 $u^2\equiv xt^2(mod m)$，又因为 t,m 互素，因此必存在一个 $t^{-1}$ 使得 $t^{-1}t\equiv 1(mod m)$ ，因此 $(ut^{-1})^2\equiv x(mod m)$，与 $QR(m,x)=0$ 矛盾，原题得证。  
### Knowledge soundness
只要 prover 同时得到 t 和 st，因为 t,m 互素，则 prover 可以计算 $s=st*t^{-1}(mod m)$ ，然后 prover 就可以欺骗 verifier.
### Zero Knowledge
1) b=0 时仅公开 t 自然是无法得知 s 的；
2) b=1 时仅公开 st，显然无法直接得知 s。又因为无法直接得知 $t^{-1}$ ，因此也无法通过 $st\*t^{-1}$ 间接计算出 s，因此是零知识的。

## 3. Self-pairing implies failure of DDH
假设 $y=y_0g$ ，则验证 $\alpha\beta g=y$ 等价于验证 $\alpha\beta g=y_0g$ ，进一步等价于验证 $\alpha\beta=y_0$ ，因此只需验证 $e(g,g)^{\alpha\beta}=e(g,g)^{y_0}$ ，等价于验证 $e(\alpha g,\beta g)=e(g,y)$ 。

## 4. BLS signature aggregation
### BLS Verify
假设 $\sigma=sH(m)$ ，若 $e(g_0,\sigma)=e(pk,H(m))$ ，则 $e(g0,H(m))^s=e(g_0,H(m))^\alpha$ ，则必有 $s=\alpha$ 。

### Multisig

$$e(g_0,\sigma)={\sum}_{i=1}^ne(g_0,\sigma_i)={\sum}_{i=1}^ne(g_0,\alpha_iH(m_i))={\sum}_{i=1}^ne(\alpha_ig_0,H(m_i))={\sum}_{i=1}^ne(pk_i,H(m_i))$$


[more reference about BLS multi signature](https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html)
