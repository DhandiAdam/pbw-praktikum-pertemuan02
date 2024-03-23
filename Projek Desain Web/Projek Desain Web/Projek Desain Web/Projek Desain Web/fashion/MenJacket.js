function masukannama() {
    if (document.forms["f"].Nama.value === "") {
      alert("Anda belum memasukkan nama");
    }
  }
  
  function masukantelpon() {
    if (document.forms["f"].Telpon.value === "") {
      alert("Anda belum memasukkan nomor telepon");
    }
  }
  
  function masukanaEmail() {
    if (document.forms["f"].email.value === "") {
      alert("Anda belum memasukkan email");
    }
  }
  
  function masukanalamat() {
    if (document.forms["f"].alamat.value === "") {
      alert("Anda belum memasukkan alamat");
    }
  }
  
  function Input() {
    var Nama = document.forms["f"].Nama.value;
    var Telpon = document.forms["f"].Telpon.value;
    var email = document.forms["f"].email.value;
    var alamat = document.forms["f"].alamat.value;
  
    var pesan = "";
  
    if (Nama === "") {
      pesan += "Anda belum memasukkan nama.<br>";
    }
    if (Telpon === "") {
      pesan += "Anda belum memasukkan nomor telepon.<br>";
    }
    if (email === "") {
      pesan += "Anda belum memasukkan email.<br>";
    }
    if (alamat === "") {
      pesan += "Anda belum memasukkan alamat.<br>";
    }
  
    if (pesan === "") {
      pesan = "Terima Kasih Telah Memesan Purple ini., Rekening BTN:547747875";
      
    }
  
    document.getElementById("resultMessage").innerHTML = pesan;
  }
  