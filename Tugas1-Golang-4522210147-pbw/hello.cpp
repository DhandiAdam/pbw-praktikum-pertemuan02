#include <iostream>
using namespace std;

typedef int angka;
typedef float pecahan;
typedef char huruf[30];

int main() {
    angka umur;
    pecahan pecah;
    huruf nama;
    huruf karakter;
	
	string pengenalan;

    cout << "Masukkan umur anda: ";
    cin >> umur;
    if(umur < 18){
		cout << "Umur Kamu Masih Anak Anak" << endl;
	}else{
		cout << "Sekarang Kamu Sudah Dewasa" << endl;
	}
	cout << "<<===========>>>" << endl;

    cout << "Masukkan bilangan pecahan: ";
    cin >> pecah;
    if (pecah == 2.5){
		cout << "Bilangan Ini yang aku cari" << endl;
	}else {
		cout << "Bukan Bilangan Itu yang aku Cari"<<endl;
	}
	
    cout << "Masukkan nama Kamu: ";
    cin >> nama;
	cout << "Nama Kamu Adalah : " << nama << endl;
	cout << "<<<<==============>>>" << endl;
	
	cout << "Masukkan nama yang ingin kamu kenal: ";
    cin >> pengenalan;
	if (pengenalan == "Dhandi"){
		cout << "Halooo Dhandi, Senang Belajar Bareng dengan anda"<<endl;
	}else if (pengenalan == "Nabila"){
		cout << "Kakaknya Cantikk bangget" << endl;
	}else {
		cout << "Boleh Kenalan tidak Nama Kamu Siapa?" << endl;
	}
	cout << "<<===========================>>" << endl;

    cout << "Masukkan satu huruf: ";
    cin >> karakter;
    cout << "Huruf anda " << karakter << endl;
	
    return 0;
}