# SpoofDPI Turkiye

Read in other Languages: [🇹🇷Turkish](https://github.com/bariiss/SpoofDPI-Turkiye), [🇬🇧English](https://github.com/bariiss/SpoofDPI-Turkiye/blob/main/_docs/README_en.md)

Spoof DPI'ın bu sürümü **Türkiye'de** kullanılmak üzere yapılandırılmıştır.

![image](https://user-images.githubusercontent.com/45588457/148035986-8b0076cc-fefb-48a1-9939-a8d9ab1d6322.png)

# Kurulum
Direkt olarak [releases](https://github.com/bariiss/SpoofDPI-Turkiye/releases) bölümünden indirebilir veya
[Buradan](https://github.com/bariiss/SpoofDPI-Turkiye/blob/main/_docs/INSTALL.md) kurulum aşamalarını takip edebilirsiniz.

# Kullanım
Programımız Türkiye'ye özel olarak konfigure edildiği için sizin için uygun sürümü direkt olarak başlatarak çalıştırabilirsiniz.

# Gelişmiş Kullanım
```
Kullanım: spoofdpi [seçenekler...]
  -addr string
        adresi dinler (varsayılan "127.0.0.1")
  -debug
        hata ayıklamayı aktif edeer
  -dns-addr string
        dns adresi (varsayılan "77.88.8.8")
  -dns-ipv4-only
        sadece sürüm 4 adreslerini dinler
  -dns-port value
        dns için port numarası (varsayılan 1253)
  -enable-doh
        'dns-over-https' aktif eder
  -pattern value
        DPI'yı yalnızca bu regex deseniyle eşleşen paketlerde atlar; birden çok kez verilebilir
  -port value
        port (varsayılan 8080)
  -silent
        başlangıçta afişi ve sunucu bilgilerini gösterme
  -system-proxy
        sistem genelinde proxy aktif et (varsayılan true)
  -timeout value
        milisaniye cinsinden zaman aşımı; verilmediğinde zaman aşımı olmaz
  -v    spoofdpi'nin sürümünü yazdırır; bu, diğer bazı ilgili bilgileri içerebilir
  -window-size value
        Parçalanmış istemci dönüşü için bayt sayısı cinsinden yığın boyutu,
        varsayılan değer DPI'ı atlamazsa daha düşük değerler deneyin;
        verilmediğinde, istemci dönüş paketi iki parça halinde gönderilecektir:
        ilk veri paketi için parçalama ve geri kalanı şeklinde
```
> Chrome tarayıcısında Hotspot Shield gibi herhangi bir vpn uzantısı kullanıyorsanız,
  Ayarlar > Eklentiler, bölümüne gidin ve onları devre dışı bırakın.

### OSX
`Spoofdpi`ı çalıştırdığınızda proxy'nizi otomatik olarak ayarlayacaktır

### Linux
`Spoofdpi`ı çalıştırın ve favori tarayıcınızı proxy seçeneği ile açın
```bash
google-chrome --proxy-server="http://127.0.0.1:8080"
```

# Nasıl Çalışır
### HTTP
 Dünyadaki çoğu web sitesi artık HTTPS'yi desteklediğinden, SpoofDPI HTTP istekleri için Derin Paket Denetimlerini atlamaz, ancak yine de tüm HTTP istekleri için proxy bağlantısı sunar.

### HTTPS
 TLS her handshake işlemini şifrelese de, İstemci dönüş paketinde alan adları hala düz metin olarak gösterilir.
 Başka bir deyişle, başka biri pakete baktığında, paketin nereye gittiğini kolayca tahmin edebilir.
 DPI işlenirken alan adı önemli bilgiler sunabilir ve aslında İstemci dönüş paketini gönderdikten hemen sonra bağlantının engellendiğini görebiliriz.
 Bunu aşmak için bazı yollar denedim ve İstemci dönüş paketini parçalara bölerek gönderdiğimizde yalnızca ilk parçanın denetlendiğini fark ettim.
 SpoofDPI'ın bunu atlamak için yaptığı şey, bir isteğin ilk 1 baytını sunucuya göndermektir,
 ve sonra geri kalanını gönder.


# Benzer Projeler
[GoodbyeDPI-Turkey](https://github.com/cagritaskn/GoodbyeDPI-Turkey) @cagritaskn
[SpoofDPI-Turkiye](https://github.com/renardev/SpoofDPI-Turkiye) @renardev
