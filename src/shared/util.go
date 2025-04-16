package shared

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"threads/src/domain"
	"unicode"
)

func GenerarUsernameDisponibleDesdeRepositorio(nombre string, repo domain.UserRepository) (string, error) {
	base := generarUsernameBase(nombre)
	if base == "@" || base == "@." {
		base = "@user"
	}

	log.Println("🔍 Base username generado:", base)

	username := base
	sufijo := 1

	for {
		log.Println("➡️ Verificando username:", username)

		u, err := repo.FindByUsername(username)
		if err != nil {
			log.Println("❌ Error al buscar username:", err)
			return "", err
		}
		if !u.Exists() {
			break
		}
		username = fmt.Sprintf("%s%d", base, sufijo)
		sufijo++
	}

	log.Println("✅ Username disponible:", username)

	return username, nil
}

func generarUsernameBase(nombre string) string {
	nombre = strings.ToLower(nombre)
	nombre = quitarTildes(nombre)

	re := regexp.MustCompile(`[^a-zA-Z\s]`)
	nombre = re.ReplaceAllString(nombre, "")

	partes := strings.Fields(nombre)

	username := "@"
	if len(partes) > 0 {
		username += partes[0]
	}
	if len(partes) > 1 {
		username += "." + partes[1]
	}

	return username
}

func quitarTildes(s string) string {
	var sb strings.Builder
	for _, r := range s {
		switch unicode.ToLower(r) {
		case 'á':
			sb.WriteRune('a')
		case 'é':
			sb.WriteRune('e')
		case 'í':
			sb.WriteRune('i')
		case 'ó':
			sb.WriteRune('o')
		case 'ú', 'ü':
			sb.WriteRune('u')
		case 'ñ':
			sb.WriteRune('n')
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
