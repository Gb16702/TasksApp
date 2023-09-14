<div class="spinner spinner-1"></div>

<style>
   /* $bg-color: #fff; //couleur de fond du conteneur;
$basic-dark-color: #212121; //couleur du spinner
$border-width: 4px; //largeur de la bordure du spinner
$basic-spinner-dimensions: 125px; //largeur et hauteur du spinner
$main-spinner-dimensions: $basic-spinner-dimensions - $border-width * 2; //largeur et hauteur du cercle le plus grand
$small-spinner-dimensions: $main-spinner-dimensions * 0.7; //largeur et hauteur du cercle le plus petit */

.spinner {
    position: relative;
    width: 32px;
    height: 32px;
}

.spinner:before,
.spinner:after {
    content: "";
    display: block;
    position: absolute;
    border-width: 2px; /* Réduisez la largeur de la bordure */
    border-style: solid;
    border-radius: 50%;
}

.spinner.spinner-1:before {
    width: calc(32px - 5px * 2);
    height: calc(32px - 5px * 2);
    border-bottom-color: black;
    border-right-color: black;
    border-top-color: rgba(0, 0, 0, 0);
    border-left-color: rgba(0, 0, 0, 0);
    top: calc(50% - (32px - 5px * 2) / 2); /* Centrage vertical */
    left: calc(50% - (32px - 5px * 2) / 2); /* Centrage horizontal */
    animation: rotate-animation 1s linear 0s infinite;
}

.spinner.spinner-1:after {
    width: calc((32px - 5px * 2) * 0.7);
    height: calc((32px - 5px * 2) * 0.7);
    border-bottom-color: black;
    border-right-color: black;
    border-top-color: rgba(0, 0, 0, 0);
    border-left-color: rgba(0, 0, 0, 0);
    top: calc(50% - ((32px - 5px * 2) * 0.7) / 2); /* Centrage vertical */
    left: calc(50% - ((32px - 5px * 2) * 0.7) / 2); /* Centrage horizontal */
    animation: anti-rotate-animation 0.85s linear 0s infinite;
}

/* Le reste de votre code reste inchangé */


@keyframes rotate-animation {
      0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(360deg);
  }
}

@keyframes anti-rotate-animation {
  0% {
    transform: rotate(0deg);
  }

  100% {
    transform: rotate(-360deg);
  }
}

</style>