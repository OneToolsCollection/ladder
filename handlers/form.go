package handlers

import "github.com/gofiber/fiber/v2"

func Form(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(html)
}

const html = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Simple HTML Form</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
</head>
    <style>
        body {
            background-color: #ffffff;
        }

        header h1 {
            text-transform: uppercase;
            font-size: 70px;
            font-weight: 600;
            color: #fdfdfe;
            text-shadow: 0px 0px 5px #7AA7D1, 0px 0px 10px #7AA7D1, 0px 0px 10px #7AA7D1,
                0px 0px 20px #7AA7D1;
        }
        .logo-title {
            font-family: 'Arial', sans-serif;
            font-size: 2rem;
            color: #fff;
            margin-bottom: 20px;
        }
        .logo {
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
		<div class="logo">
			<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="250" height="250" viewBox="0 0 512 512">
				<path fill="#7AA7D1" d="M262.074 485.246C254.809 485.265 247.407 485.534 240.165 484.99L226.178 483.306C119.737 468.826 34.1354 383.43 25.3176 274.714C24.3655 262.975 23.5876 253.161 24.3295 241.148C31.4284 126.212 123.985 31.919 238.633 24.1259L250.022 23.8366C258.02 23.8001 266.212 23.491 274.183 24.1306C320.519 27.8489 366.348 45.9743 402.232 75.4548L416.996 88.2751C444.342 114.373 464.257 146.819 475.911 182.72L480.415 197.211C486.174 219.054 488.67 242.773 487.436 265.259L486.416 275.75C478.783 352.041 436.405 418.1 369.36 455.394L355.463 462.875C326.247 477.031 294.517 484.631 262.074 485.246ZM253.547 72.4475C161.905 73.0454 83.5901 144.289 73.0095 234.5C69.9101 260.926 74.7763 292.594 83.9003 317.156C104.53 372.691 153.9 416.616 211.281 430.903C226.663 434.733 242.223 436.307 258.044 436.227C353.394 435.507 430.296 361.835 438.445 267.978C439.794 252.442 438.591 236.759 435.59 221.5C419.554 139.955 353.067 79.4187 269.856 72.7052C264.479 72.2714 258.981 72.423 253.586 72.4127L253.547 72.4475Z"/>
				<path fill="#7AA7D1" d="M153.196 310.121L133.153 285.021C140.83 283.798 148.978 285.092 156.741 284.353L156.637 277.725L124.406 278.002C123.298 277.325 122.856 276.187 122.058 275.193L116.089 267.862C110.469 260.975 103.827 254.843 98.6026 247.669C103.918 246.839 105.248 246.537 111.14 246.523L129.093 246.327C130.152 238.785 128.62 240.843 122.138 240.758C111.929 240.623 110.659 242.014 105.004 234.661L97.9953 225.654C94.8172 221.729 91.2219 218.104 88.2631 214.005C84.1351 208.286 90.1658 209.504 94.601 209.489L236.752 209.545C257.761 209.569 268.184 211.009 285.766 221.678L285.835 206.051C285.837 197.542 286.201 189.141 284.549 180.748C280.22 158.757 260.541 143.877 240.897 135.739C238.055 134.561 232.259 133.654 235.575 129.851C244.784 119.288 263.68 111.99 277.085 111.105C288.697 109.828 301.096 113.537 311.75 117.703C360.649 136.827 393.225 183.042 398.561 234.866C402.204 270.253 391.733 308.356 367.999 335.1C332.832 374.727 269.877 384.883 223.294 360.397C206.156 351.388 183.673 333.299 175.08 316.6C173.511 313.551 174.005 313.555 170.443 313.52L160.641 313.449C158.957 313.435 156.263 314.031 155.122 312.487L153.196 310.121Z"/>
			</svg>
		</div>
        <header>
            <h1 class="center-align logo-title">ladddddddder</h1>
        </header>
        <form id="inputForm" class="col s12" method="get">
            <div class="row">
                <div class="input-field col s10">
                    <input type="text" id="inputField" name="inputField" class="validate" required>
                    <label for="inputField">URL</label>
                </div>
				<!--
                <div class="input-field col s2">
                    <button class="btn waves-effect waves-light" type="submit" name="action">Submit
                        <i class="material-icons right">go</i>
                    </button>
                </div>
				-->
            </div>
        </form>
    </div>

    <script src="https://code.jquery.com/jquery-3.6.0.min.js"></script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>
    <script>
        document.addEventListener('DOMContentLoaded', function() {
            M.AutoInit();
        });
        document.getElementById('inputForm').addEventListener('submit', function (e) {
            e.preventDefault();
            const inputValue = document.getElementById('inputField').value;
            window.location.href = '/' + inputValue;
            return false;
        });
    </script>
</body>
</html>
`
