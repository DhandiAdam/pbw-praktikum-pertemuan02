
// Tailwind Css
const navTriggerBtn = document.querySelector('#nav_trigger_btn');

const navMenu = document.querySelector('#nav_menu');

navTriggerBtn.addEventListener('click', () => {
  navMenu.classList.toggle('nav-is-open');
});

// ReadText
function ReadText() {
  var moreText = document.getElementById("moreText");
  moreText.classList.toggle("hidden");

  var btnText = moreText.classList.contains("hidden") ? "More Detail" : "Less More";
  document.querySelector(".main_btn").innerText = btnText;
}
// Orderan
function Orderan() {
  var moreText = document.getElementById("orderan");
  moreText.classList.toggle("hidden");

  var btnText = moreText.classList.contains("hidden") ? "More Detail" : "Less More";
  document.querySelector(".Orderan").innerText = btnText;
}

// Reservasi
function Reservasi() {
  var moreText = document.getElementById("reservasi");
  moreText.classList.toggle("hidden");

  var btnText = moreText.classList.contains("hidden") ? "More Detail" : "Less More";
  document.querySelector(".Reservasi").innerText = btnText;
}


// Redirect to WhattsApp

document.getElementById("waButton").addEventListener("click", redirectToWhatsApp);

function redirectToWhatsApp() {
    var phoneNumber = '6281282135908'; 
    var message = 'Hello, saya ingin melakukan reservasi di restoran Anda. Berikut adalah detail reservasi saya:\n\nNama: \nTanggal: \nWaktu: \nJumlah Orang: \n\nTerimakasih!'; 
    var url = `https://wa.me/${phoneNumber}?text=${encodeURIComponent(message)}`;
    window.open(url, '_blank');
}   

function TempatMewah1(){
  var url = `https://www.google.com/maps/place/Jl.+Kemang+Raya+No.29,+RT.1%2FRW.7,+Bangka,+Kec.+Mampang+Prpt.,+Kota+Jakarta+Selatan,+Daerah+Khusus+Ibukota+Jakarta+12720/@-6.2550212,106.8096315,17z/data=!3m1!4b1!4m6!3m5!1s0x2e69f1790d42d63f:0x4ad34399996fdc23!8m2!3d-6.2550212!4d106.8096315!16s%2Fg%2F11c1ycfnxc?entry=ttu`;
  window.open(url, '_blank');
}
function TempatMewah2(){
  var url = `https://www.google.com/maps/place/Sofia+at+The+Gunawarman/@-6.2439483,106.7720042,13z/data=!4m6!3m5!1s0x2e69f15b16b9269d:0x576ab165b739f1f7!8m2!3d-6.2315139!4d106.8077447!16s%2Fg%2F11cn8t175f?entry=ttu`;
  window.open(url, '_blank');
}
function TempatMewah3(){
var url =`https://www.google.com/maps/place/The+Energy+Building/@-6.2262521,106.8037811,17z/data=!3m1!4b1!4m6!3m5!1s0x2e69f57d5013eb65:0xb2d707d4655b2d8!8m2!3d-6.2262521!4d106.806356!16s%2Fm%2F0wfb551?entry=ttu`;
  window.open(url, '_blank');
}
function Memesan() {
  var phoneNumber = '6281282135908'; 
  var message = 'Hello, saya ingin Memesan Paket yang ada diwebsite, bagaimana yah carannya?'; 
  var url = `https://wa.me/${phoneNumber}?text=${encodeURIComponent(message)}`;
  window.open(url, '_blank');
}  

const swiper = new Swiper('.swiper',{
  loop: true,
  pagination: {
    el: '.swiper-pagination',
    clickable: true,
  },
  slidesPerView: 3,
  spaceBetween: 20,

  breakpoints:{
    320:{
      slidesPerView: 1,
      
    },
    960: {
      slidesPerView: 2,
    },
    1200: {
      slidesPerView: 3,
    },
  },

});

const sr = ScrollReveal({
  origin: 'bottom',
  distance: '60px',
  duration: 3000,
  // reset: true,
  delay: 600,
})

// Hero 
sr.reveal('.hero__text',{origin: 'top'});

// Step
sr.reveal('.steps__step',{distance: '100px', interval:100});
// About
sr.reveal('.about__text',{origin: 'left'});
sr.reveal('.about__img',{origin: 'right', delay: 800});
// Favorite
sr.reveal('.Favorite__bg',{delay: 800});
sr.reveal('.Favorite__title');
sr.reveal('.Favorite__slider',{delay: 1000});

// Reservasi
sr.reveal('.Reservasi__title');
sr.reveal('.Reservasi__subtitle', {delay:800});
sr.reveal('.Makanan__grid', {delay:1000});

// Stats
sr.reveal('.stats');
sr.reveal('.stats__item', {
  distance: '100px',
  interval: 100,
});
// Tempat Reservasi
sr.reveal('.Tempat__title');
sr.reveal('.Tempat__subtitle', {delay: 800});
sr.reveal('.Tempat__item', {
  distance: '100px',
  interval: 100,
  delay: 1000,
});

// Contact 

sr.reveal('.Contact__container');
sr.reveal('.Contact__text', {delay:800});

// Footer
sr.reveal('.footer__item', {
  distance: '100px',
  interval: 100,
  
});
sr.reveal('.footer__copyright');

