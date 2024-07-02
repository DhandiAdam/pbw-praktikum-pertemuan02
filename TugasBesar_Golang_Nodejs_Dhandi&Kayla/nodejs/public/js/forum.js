document
            .getElementById("billing-form")
            .addEventListener("submit", async (e) => {
                e.preventDefault();

                const name = document.getElementById("name").value;
                const address = document.getElementById("address").value;
                const phone = document.getElementById("phone").value;
                const items = JSON.parse(localStorage.getItem("cartItems"));

                const response = await fetch("/submit-billing", {
                    method: "POST",
                    headers: { "Content-Type": "application/json" },
                    body: JSON.stringify({ name, address, phone, items }),
                });

                const result = await response.json();

                if (result.success) {
                    // Redirect to Stripe checkout
                    const checkoutResponse = await fetch("/stripe-checkout", {
                        method: "POST",
                        headers: { "Content-Type": "application/json" },
                        body: JSON.stringify({ items }),
                    });

                    const checkoutResult = await checkoutResponse.json();
                    window.location.href = checkoutResult.url;
                    clearCart();
                } else {
                    console.error("Error submitting billing info:", result.error);
                }
            });

        function clearCart() {
            localStorage.removeItem('cartItems');
            const cartContent = document.querySelector('.cart-content');
            while (cartContent.firstChild) {
                cartContent.removeChild(cartContent.firstChild);
            }
        }