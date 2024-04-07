async function checkID() {
    try {
        let certificate_id = document.getElementById("certificate_id");
        let inputResponse = document.getElementById("id-sertif-input-response");

        if (!certificate_id.value) {
            inputResponse.innerHTML = "<div>Tolong masukan ID anda</div>";

            return Swal.fire({
                icon: "error",
                title: "Gagal",
                text: "Tolong isi field dengan benar!",
                buttonsStyling: false,
                confirmButtonText: "Ok",
                customClass: {
                    confirmButton: "btn btn-danger",
                },
            });
        }

        const url =
            BASE_URL + "/api/certificate/validate/id/" + certificate_id.value;

        const validate = await fetch(url);

        const resLogin = await validate.json();

        if (validate.status === 200) {
            Swal.fire({
                text: "Sertifikat anda valid!",
                icon: "success",
                buttonsStyling: false,
                confirmButtonText: "Selanjutnya",
                customClass: {
                    confirmButton: "btn btn-primary",
                },
            }).then((result) => {
                if (result.isConfirmed) {
                    // window.location.href = BASE_URL + "/dashboard";
                }
            });
        } else {
            return Swal.fire({
                icon: "error",
                title: "Login Gagal",
                text: resLogin.Val.message,
            });
        }
    } catch (error) {
        return Swal.fire({
            icon: "error",
            title: "Login Gagal",
            text: error,
        });
    }
}
