package teleinfo

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	tagErrorType = "error_type"
)

var (
	// Exporter metrics
	frameReadCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "teleinfo_frames_read_total",
		Help: "The total number of read Teleinfo frames",
	})
	frameReadErrorCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teleinfo_frames_read_errors_total",
			Help: "The total number of frame read errors",
		},
		[]string{tagErrorType},
	)

	frameDecodedCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "teleinfo_frames_decoded_total",
		Help: "The total number of decoded frames",
	})
	frameDecodeErrorCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "teleinfo_frames_decode_errors_total",
			Help: "The total number of frame decoding errors",
		},
		[]string{tagErrorType},
	)

	// Teleinfo metrics
	// "ADCO": "004322423452",  Adresse du concentrateur de téléreport // no metric
	// "PPOT": "00",           Présence des potentiels // no metric

	// "HCHC": "000932141",     Heures Creuses (Wh)
	teleinfoHeureCreuseGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "teleinfo_heures_creuses_total",
			Help: "Total des Heures Creuses en Wh",
		})
	// "HCHP": "002663019",     Heures Pleines (Wh)
	teleinfoHeurePleinesGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "teleinfo_heures_pleines_total",
			Help: "Total des Heures Pleines en Wh",
		})
		// "BASE": "002663019",     Heures de base (Wh)
		teleinfoBaseGauge = promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "teleinfo_base_total",
				Help: "Total tarif de base en Wh",
			})
	// "HHPHC": "A",            Horaire Heures Pleines Heures Creuses
	teleinfoProgrammationHeuresPleinesHeuresCreusesGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "teleinfo_programmation_heures_pleines_heures_creuses",
			Help: "Programmation du compteur pour l'horaire heurespleines/heurescreuses",
		},
		[]string{"programme"},
	)
	// "IINST": "003",         Intensité Instantanée par phase (A)
	// "IINST1": "000",        Intensité Instantanée par phase (A)
	// "IINST2": "006",        Intensité Instantanée par phase (A)
	// "IINST3": "001",        Intensité Instantanée par phase (A)
	teleinfoIntensiteInstantaneeGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "teleinfo_instensite_instantanee",
			Help: "Intensite Instantanée par phase en A",
		},
		[]string{"phase"},
	)
	// "IMAX1": "060",         Intensité maximale par phase (A)
	// "IMAX2": "060",         Intensité maximale par phase (A)
	// "IMAX3": "060",         Intensité maximale par phase (A)
	teleinfoIntensiteMaximaleGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "teleinfo_instensite_maximale",
			Help: "Intensite Maximale par phase en A",
		},
		[]string{"phase"},
	)
	// "ISOUSC": "15",         intensité souscrite (A)
	teleinfoIntensiteSouscriteGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "teleinfo_instensite_souscrite",
			Help: "Intensite Souscrite en A",
		},
	)
	// "MOTDETAT": "000000",   Mot d'Etat du compteur
	teleinfoModeEtatCompteurGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "teleinfo_mode_etat_compteur",
			Help: "Mot d'Etat du compteur",
		},
	)
	// "OPTARIF": "HC..",      Option tarifaire choisie
	teleinfoOptionTarifaireChoisieGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "teleinfo_option_tarifaire_choisie",
			Help: "Option tarifaire choisie",
		},
		[]string{"tarif"},
	)
	// "PAPP": "01790",        Puissance apparente triphasée (VA)
	teleinfoPuissanceApparenteTriphaseGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "teleinfo_puissance_apparente_triphase",
			Help: "Puissance apparente triphasée en VA",
		},
	)
	// "PMAX": "07250",        Puissance maximale triphasée atteinte (W)
	teleinfoPuissanceMaximaleTriphaseGauge = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "teleinfo_puissance_maximale_triphase",
			Help: "Puissance maximale triphasée atteinte en W",
		},
	)
	// "PTEC": "HP.."          période tarifaire en cours
	teleinfoPeriodeTarifaireEnCoursGauge = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "teleinfo_periode_tarifaire_en_cours",
			Help: "Periode tarifaire en cours",
		},
		[]string{"tarif"},
	)
)

func incrementErrorCounter(counter *prometheus.CounterVec, errorType string) {
	counter.With(prometheus.Labels{tagErrorType: errorType}).Inc()
}
